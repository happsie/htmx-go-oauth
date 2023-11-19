package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/happise/pixelwars/container"
	"github.com/happise/pixelwars/model"
)

type JWTService interface {
	CreateJWT(user model.TwitchUser, expiry int64) (string, error)
}

type jwtService struct {
	container container.Container
}

func NewJWTService(container container.Container) JWTService {
	return &jwtService{
		container: container,
	}
}

func (js jwtService) CreateJWT(user model.TwitchUser, expiry int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      user.ID,
		"displayName": user.DisplayName,
		"exp":         expiry,
	})
	tokenString, err := token.SignedString([]byte(js.container.GetConfig().JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
