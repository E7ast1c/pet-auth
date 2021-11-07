package database

import (
	"context"

	"pet-auth/internal/config"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"go.uber.org/zap"
)

func NewConnect(ctx context.Context, config *config.AppConfig) (*pgx.Conn, error) {
	pgConf, err := pgx.ParseConfig(config.DatabaseConfig.URI)
	if err != nil {
		return nil, err
	}
	pgConf.Logger = zapadapter.NewLogger(zap.L())

	conn, err := pgx.ConnectConfig(ctx, pgConf)
	if err != nil {
		return conn, nil
	}

	return conn, nil
}
