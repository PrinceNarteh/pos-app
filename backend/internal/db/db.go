// Package db provides database initialization and connection management.
package db

import (
	"context"
	"time"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), config.Env.DB.URI)
	if err != nil {
		return nil, err
	}

	maxIdleTime, err := time.ParseDuration(config.Env.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.Config().MaxConns = config.Env.DB.MaxOpenConns
	db.Config().MinConns = config.Env.DB.MinOpenConns
	db.Config().MaxConnIdleTime = maxIdleTime

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
