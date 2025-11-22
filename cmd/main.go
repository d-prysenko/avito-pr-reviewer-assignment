package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"revass/internal/config"
	"revass/internal/router"
	"revass/internal/services/logger"
	"revass/internal/storage/postgersql"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg)
	slog.SetDefault(log)

	database := postgersql.Connect(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Schema)

	go func() {
		<-ctx.Done()
		database.Close()
	}()

	r := router.New()
	addr := fmt.Sprintf(":%d", cfg.Port)

	fmt.Printf("Starting server on %s\n", addr)

	http.ListenAndServe(addr, r)
}
