package server

import (
	"fmt"
	"net/http"

	"pet-auth/internal/auth"
	"pet-auth/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Error   string
	Payload string
}

func (w WebServer) Login(c *gin.Context) {

}

func (w WebServer) Logout(c *gin.Context) {

}

func (w WebServer) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error:   fmt.Sprintf("error on deserialization: %s", err.Error()),
			Payload: "",
		})
		return
	}

	pHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Error:   fmt.Sprintf("password hashing failed: %s", err.Error()),
			Payload: "",
		})
		return
	}

	user.Password = string(pHash)
	rUser, err := w.repo.Register(&user, w.config.JWTSettings.SignAlgorithm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Error:   fmt.Sprintf("error on db register: %s", err.Error()),
			Payload: "",
		})
		return
	}

	jwt, err := auth.NewJWTAuth(*w.config, *w.repo).CreateJWT(rUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Error:   fmt.Sprintf("error on db register: %s", err.Error()),
			Payload: "",
		})
		return
	}

	c.JSON(http.StatusOK, struct {
		AccessToken string `json:"access_token"`
		Name        string `json:"name"`
		Email       string `json:"email"`
	}{jwt, rUser.Name, rUser.Email})
}
