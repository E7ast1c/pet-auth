package repository

import (
	"context"

	"pet-auth/internal/database"
	"pet-auth/internal/models"
)

type Auth interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
	Logout()
}

type Repository struct {
	pgDb database.PGDB
	ctx  context.Context
}

func NewRepository(ctx context.Context, pgDb database.PGDB) *Repository {
	return &Repository{pgDb: pgDb, ctx: ctx}
}
