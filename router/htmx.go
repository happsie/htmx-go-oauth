package router

import (
	"fmt"
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/handlers"
	"github.com/happsie/gohtmx/repository"
	"github.com/happsie/gohtmx/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterPaths(e *echo.Echo, container container.Container) {
	userRepo := repository.NewUserRepository(container)
	// TODO: move to view handler both index and home
	e.GET("/", func(c echo.Context) error {
		_, err := utils.GetAuthInfo(c, container.GetConfig())
		if err != nil {
			return c.Render(200, "index.html", nil)
		}
		return c.Redirect(http.StatusSeeOther, "/home")
	})
	e.GET("/home", func(c echo.Context) error {
		auth, err := utils.GetAuthInfo(c, container.GetConfig())
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}
		user, err := userRepo.Get(auth.UserId)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.Render(200, "home.html", user)
	}, echojwt.WithConfig(GetJwtConfig(container)))

	// Command handling
	commandHandler := handlers.NewCommandHandler(container)
	commands := e.Group("/commands", echojwt.WithConfig(GetJwtConfig(container)))
	commands.GET("", commandHandler.List)
	commands.GET("/modal", commandHandler.Modal)
	commands.POST("/add", func(c echo.Context) error {
		fmt.Println("form value: " + c.FormValue("command"))
		return commandHandler.List(c)
	})
}
