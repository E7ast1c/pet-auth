package server

import (
	"net/http"

	"pet-auth/internal/config"
	"pet-auth/internal/middleware"
	"pet-auth/internal/models"
	"pet-auth/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Server interface {
	SetHandlers()
	Run()

	Login(c *gin.Context)
	Logout(c *gin.Context)
	Register(c *gin.Context)
}

func NewWebServer(db *pgx.Conn, config *config.AppConfig, repo *repository.Repository) *WebServer {
	return &WebServer{
		router: gin.Default(),
		db:     db,
		config: config,
		repo: repo,
	}
}

type WebServer struct {
	router *gin.Engine
	db     *pgx.Conn
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

func (w WebServer) Login(c *gin.Context) {

}

func (w WebServer) Logout(c *gin.Context) {

}

func (w WebServer) Register(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error on deserialization": err.Error()})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"password encryption failed": err.Error()})
		return
	}

	zap.L().Info(string(pass))

	registeredUser, err := w.repo.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error on db register": err.Error()})
		return
	}

	zap.S().Info(registeredUser)
}
