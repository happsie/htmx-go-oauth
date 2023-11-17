package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/happise/pixelwars/config"
	"github.com/happise/pixelwars/model"
	"github.com/labstack/echo/v4"
)

func GetAuthInfo(c echo.Context, config config.Config) (model.JwtInfo, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return model.JwtInfo{}, err
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	if err != nil {
		return model.JwtInfo{}, err
	}
	jwtInfo := model.JwtInfo{
		UserId:   fmt.Sprintf("%v", claims["userId"]),
		Username: fmt.Sprintf("%v", claims["displayName"]),
		Exp:      0,
	}
	return jwtInfo, nil
}
