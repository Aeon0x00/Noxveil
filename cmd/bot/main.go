package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aeon0x00/Noxveil/internal/bot"
	"github.com/Aeon0x00/Noxveil/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	//"github.com/k0kubun/pp"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")
	tgBot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Ошибка создания postgres pool")
	}

	if err := pool.Ping(ctx); err != nil {
		log.Println("Ошибка пинга к базе данных")
	}

	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	subRepo := repository.NewSubRepo(pool)

	updates, _ := tgBot.UpdatesViaLongPolling(ctx, nil)

	bh, _ := th.NewBotHandler(tgBot, updates)

	defer func() { _ = bh.Stop() }()
 
	h := bot.NewHandler(userRepo,subRepo)
	h.Register(bh)
	
	_ = bh.Start()
}
