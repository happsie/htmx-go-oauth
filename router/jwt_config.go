package router

import (
	"github.com/happise/pixelwars/container"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetJwtConfig(container container.Container) echojwt.Config {
	return echojwt.Config{
		TokenLookup: "cookie:token",
		SigningKey:  []byte(container.GetConfig().JWT.Secret),
		ErrorHandler: func(c echo.Context, err error) error {
			cookie := new(http.Cookie)
			cookie.Name = "token"
			cookie.Value = ""
			cookie.Expires = time.Now().AddDate(0, 0, -1)
			cookie.Path = "/"
			cookie.HttpOnly = true
			c.SetCookie(cookie)
			return c.Redirect(http.StatusSeeOther, "/")
		},
	}
}
