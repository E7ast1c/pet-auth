package config

import (
	"pet-auth/pkg/dbconv"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type AppConfig struct {
	DatabaseConfig struct {
		Port     uint32 `env:"DB_PORT" env-required:""`
		Host     string `env:"DB_HOST" env-required:""`
		User     string `env:"DB_NAME" env-required:""`
		Dbname   string `env:"DB_USER" env-required:""`
		Password string `env:"DB_PASSWORD" env-required:""`
		URI      string
	}
	WebServerConfig struct {
		Port string `env:"WEB_PORT" env-required:""`
	}
	JWTSettings struct {
		Secret        string `env:"SIGN_SECRET" env-default:"sign-secret"`
		SignAlgorithm string `env:"SIGN_ALGORITHM" env-default:"HS256"`
		Issuer        string `env:"SIGN_ISSUER" env-default:"pet-auth"`
	}
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

const defaultEnvFile = ".env"

func (a *AppConfig) Load() *AppConfig {
	err := cleanenv.ReadConfig(defaultEnvFile, a)
	if err != nil {
		zap.S().Fatal("", err)
	}
	a.DatabaseConfig.URI = dbconv.PGURLConv(a.DatabaseConfig.Port, a.DatabaseConfig.Host,
		a.DatabaseConfig.User, a.DatabaseConfig.Dbname, a.DatabaseConfig.Password)
	return a
}
