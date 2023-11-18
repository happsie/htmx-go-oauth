package router

import (
	"github.com/happise/pixelwars/container"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func GetJwtConfig(container container.Container) echojwt.Config {
	return echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  []byte(container.GetConfig().JWT.Secret),
	}
}
