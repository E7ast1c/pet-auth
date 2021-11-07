package main

import (
	"context"
	"log"

	"pet-auth/internal/container"
	"pet-auth/internal/server"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger() (*zap.Logger, error) {
	cnfg := zap.NewDevelopmentConfig()
	cnfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cnfg.EncoderConfig.TimeKey = "timestamp"
	cnfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cnfg.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func main() {
	zapLog, err := setupLogger()
	if err != nil {
		log.Fatal(err)
	}

	err = container.GetContainer().Invoke(func(server *server.WebServer) {
		server.SetHandlers()
		server.Run()
	})
	if err != nil {
		zap.S().Fatal(err)
	}

	shutdown()

	defer func(loggerMgr *zap.Logger) {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(zapLog)
}

func shutdown() {
	c := container.GetContainer()
	err := c.Invoke(func(ctx context.Context, conn *pgx.Conn) {
		err := conn.Close(ctx)
		zap.S().Info("conn closed")
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}