package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createTable(db *pgxpool.Pool) error {
	ctx := context.Background()
	_, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
				id UUID PRIMARY KEY,
				first_name TEXT NOT NULL,
				last_name TEXT NOT NULL,
				biography TEXT NOT NULL
			);`)

	if err != nil {
		return err
	}

	return nil
}
