package migrations

import (
	"context"

	"pet-auth/internal/config"
	"pet-auth/internal/container"
	"pet-auth/internal/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migration run go-migration
func Migration() error {
	c := container.GetContainer()
	var dbUrl, dbName string
	var db database.PGDB

	err := c.Invoke(func(ctx context.Context, config *config.AppConfig, PgDb database.PGDB) {
		dbUrl, dbName = config.DatabaseConfig.URI, config.DatabaseConfig.Dbname
		db = PgDb
	})
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://../pet-auth/internal/migrations", dbName, driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	return nil
}
