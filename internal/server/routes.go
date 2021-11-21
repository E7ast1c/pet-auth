package server

import (
	"pet-auth/internal/config"
	"pet-auth/internal/database"
	"pet-auth/internal/middleware"
	"pet-auth/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Server interface {
	SetHandlers()
	Run()

	Login(c *gin.Context)
	Logout(c *gin.Context)
	Register(c *gin.Context, signAlgo string)
}

func NewWebServer(db database.PGDB, config *config.AppConfig, repo *repository.Repository) *WebServer {
	return &WebServer{
		router: gin.Default(),
		pgDb:   &db,
		config: config,
		repo:   repo,
	}
}

type WebServer struct {
	router *gin.Engine
	pgDb   *database.PGDB
	config *config.AppConfig
	repo   *repository.Repository
}

func (w *WebServer) SetHandlers() {
	if _, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		zap.S().Fatalf("Gin binding.Validator.Engine is not *validator.Validate")
	}

	w.router.Use(middleware.Cors)

	w.router.POST("register", w.Register)
	w.router.POST("login", w.Login)

	auth := w.router.Group("authorize")
	{
		auth.POST("logout", w.Logout)
	}
}

func (w WebServer) Run() {
	err := w.router.Run(":" + w.config.WebServerConfig.Port)
	if err != nil {
		zap.S().Fatalf("run web server failed %s", err)
	}
}
