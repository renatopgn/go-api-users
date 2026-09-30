package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/renatopgn/Go/api"
)

func main() {

	if err := run(); err != nil {
		slog.Error("fail to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	ctx := context.Background()
	url := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return err
	}

	if err := createTable(db); err != nil {
		slog.Error("error to crate table users", "error", err)
		return err
	}

	handler := api.NewHandler(db)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
