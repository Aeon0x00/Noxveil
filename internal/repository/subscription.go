package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SubRepo struct {
	pool *pgxpool.Pool
}

func NewSubRepo (pool *pgxpool.Pool) *SubRepo {
	return &SubRepo{
		pool: pool,
	}
}

func (s *SubRepo) CheckOnActive(ctx context.Context, ID int64) (bool, error){
	var isActive bool
	err := s.pool.QueryRow(ctx, `SELECT EXISTS 
			(SELECT 1 FROM subscriptions 
			WHERE user_id = $1 
				AND is_active 
				AND (expiry_time is NULL OR expiry_time > NOW()))
				`,ID).Scan(&isActive)
	if err != nil {
		// logger
		return isActive, err
	}
	return isActive, nil
} 

func (s *SubRepo) CreateSub(ctx context.Context, ID string, userID int64, subID string,limitDevice int, totalGB int) error {
	_, err := s.pool.Exec(ctx,"INSERT INTO subscriptions (id,user_id,sub_id,limit_device,total_bytes) VALUES ($1,$2,$3,$4,$5) ",ID,userID,subID,limitDevice,totalGB*1024*1024*1024) 
	if err != nil {
		log.Println("Ошибка", err)
		return err
	}
	log.Println("Успешно")
	return nil
}