package container

import (
	"context"
	"sync"

	"pet-auth/internal/config"
	"pet-auth/internal/database"
	"pet-auth/internal/repository"
	"pet-auth/internal/server"

	"go.uber.org/dig"
	"go.uber.org/zap"
)

type Container struct {
	container *dig.Container
}

var (
	singleton *Container
	once      sync.Once
)

func buildContainer() *dig.Container {
	ctx := context.Background()
	container := dig.New()

	// web server
	err := container.Provide(server.NewWebServer)
	if err != nil {
		zap.S().Fatal(err)
	}

	// context
	err = container.Provide(func() context.Context { return ctx })
	if err != nil {
		zap.S().Fatal(err)
	}

	// configuration
	err = container.Provide(config.NewAppConfig().Load)
	if err != nil {
		zap.S().Fatal(err)
	}

	// database
	err = container.Provide(database.NewConnect)
	if err != nil {
		zap.S().Fatal(err)
	}

	// repository
	err = container.Provide(repository.NewRepository)
	if err != nil {
		zap.S().Fatal(err)
	}
	return container
}

func GetContainer() *dig.Container {
	once.Do(func() {
		singleton = &Container{buildContainer()}
	})

	return singleton.container
}