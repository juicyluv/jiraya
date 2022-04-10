package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/config"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/logging"
	"github.com/juicyluv/jiraya/internal/jiraya/storage"
	"go.uber.org/zap"
	"time"
)

type db struct {
	*pgxpool.Pool
	logger *zap.Logger
}

func New(cfg *config.Config) (storage.Storage, error) {
	pgxPoolConfig, err := pgxpool.ParseConfig(cfg.DB.DSN)

	if err != nil {
		return nil, fmt.Errorf("cannot parse postgres config: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.DB.ConnectionTimeout)*time.Second)
	defer cancel()

	pool, err := pgxpool.ConnectConfig(ctx, pgxPoolConfig)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres: %v\n", err)
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil, fmt.Errorf("cannot ping postgres: %v\n", err)
	}

	return &db{pool, logging.Get()}, nil
}
