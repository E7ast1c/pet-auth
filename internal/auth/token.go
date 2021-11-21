package auth

import (
	"errors"
	"strconv"
	"time"

	"pet-auth/internal/models"

	"github.com/dgrijalva/jwt-go"
)

const (
	oneDay = time.Hour * 24
	// RS256 sign method extend CreateJWT func
	RS256 = "RS256"
	// HS256 sign method extend CreateJWT func
	HS256 = "HS256"
)

type userToken struct {
	models.User
	*jwt.StandardClaims
}

func (a *JWTAuth) CreateJWT(user *models.User) (string, error) {
	now := time.Now()
	expire := now.Add(oneDay)

	tkn := userToken{User: models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, StandardClaims: &jwt.StandardClaims{
		Audience:  "[]",
		ExpiresAt: expire.Unix(),
		Id:        strconv.FormatInt(user.Id, 10),
		IssuedAt:  now.Unix(),
		Issuer:    a.Config.JWTSettings.Issuer,
		NotBefore: now.Unix(),
		Subject:   "Json Web Token",
	}}

	token := &jwt.Token{}
	switch a.Config.JWTSettings.SignAlgorithm {
	case RS256:
		token = jwt.NewWithClaims(jwt.SigningMethodRS256, tkn)
	case HS256:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, tkn)
	default:
		return "", errors.New("define sign method failed")
	}

	return token.SignedString([]byte(a.Config.JWTSettings.Secret))
}
