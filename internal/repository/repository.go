package repository

import (
	"pet-auth/internal/models"

	"github.com/jackc/pgx/v4"
)

type Auth interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
	Logout()
}

type Repository struct{
	db *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{conn}
}
