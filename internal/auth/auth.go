package auth

import (
	"pet-auth/internal/config"
	"pet-auth/internal/repository"
)

type JWTAuth struct {
	Repo   repository.Repository
	Config config.AppConfig
}

func NewJWTAuth(config config.AppConfig, repo repository.Repository) *JWTAuth {
	auth := &JWTAuth{
		Repo:   repo,
		Config: config,
	}

	return auth
}


