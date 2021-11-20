package database

import (
	"context"
	"database/sql"

	"pet-auth/internal/config"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

type PGDB struct {
	Conn *sql.Conn
	DB *sql.DB
}

func NewConnect(ctx context.Context, config *config.AppConfig) (PGDB, error) {
	connConfig, err := pgx.ParseConfig(config.DatabaseConfig.URI)
	if err != nil {
		return PGDB{}, err
	}

	connConfig.Logger = zapadapter.NewLogger(zap.L())
	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return PGDB{}, err
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		return PGDB{}, err
	}

	return PGDB{
		Conn: conn,
		DB:   db,
	}, nil
}
