package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/happsie/gohtmx/config"
	"github.com/happsie/gohtmx/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
	exp := claims["exp"].(float64)
	jwtInfo := model.JwtInfo{
		UserId:   fmt.Sprintf("%v", claims["userId"]),
		Username: fmt.Sprintf("%v", claims["displayName"]),
		Exp:      int64(exp),
	}
	return jwtInfo, nil
}

func SetAuthCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(5 * time.Minute) // TODO: Set time to oauth exp
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func InvalidateAuthCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().AddDate(0, 0, -30)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
