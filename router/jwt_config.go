package router

import (
	"github.com/happise/pixelwars/container"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"slices"
)

func GetJwtConfig(container container.Container) echojwt.Config {
	jwtWhitelist := []string{"/api/auth/login", "/api/auth/callback", "", "/", "/css*", "/images*"}
	return echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  []byte(container.GetConfig().JWT.Secret),
		Skipper: func(c echo.Context) bool {
			return slices.Contains(jwtWhitelist, c.Path())
		},
	}
}
