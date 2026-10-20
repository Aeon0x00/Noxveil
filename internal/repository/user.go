package repository

import (
	"context"

	"github.com/Aeon0x00/Noxveil/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)
type UserRepository struct {
	pool *pgxpool.Pool
	//log *zap......
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error){
	var user domain.User
	user.TelegramID = telegramID
	err := r.pool.QueryRow(ctx,"SELECT id, balance, created_at FROM users WHERE telegram_id = $1;",telegramID).Scan(&user.ID,&user.Balance,&user.CreatedAt)
	if err != nil {
		// logger будет обрабатывать
		return nil, err
	}
	return &user, nil
}
 
func (r *UserRepository) CreateUser(ctx context.Context, telegramID int64) (*domain.User, error) {
	var user domain.User
	user.TelegramID = telegramID
	err := r.pool.QueryRow(ctx,"INSERT INTO users (telegram_id) VALUES ($1) RETURNING id, balance, created_at;",telegramID).Scan(&user.ID,&user.Balance,&user.CreatedAt)
	if err != nil {
		// logger будет обрабатывать
		return nil, err
	}
	return  &user, nil
}