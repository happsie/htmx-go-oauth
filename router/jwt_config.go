package router

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetJwtConfig(container container.Container) echojwt.Config {
	return echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  []byte(container.GetConfig().JWT.Secret),
		ErrorHandler: func(c echo.Context, err error) error {
			utils.InvalidateAuthCookie(c)
			return c.Redirect(http.StatusSeeOther, "/")
		},
	}
}
