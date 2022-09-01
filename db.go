package main

import (
	// adding this here to initialize the pg driver in pgx
	"context"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDB(ctx context.Context, cfg cfg) (*pgx.Conn, error) {
	return pgx.Connect(ctx, cfg.DatabaseURL)
}
