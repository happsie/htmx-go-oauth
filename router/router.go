package router

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/handlers"
	"github.com/happsie/gohtmx/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func Init(container container.Container) *echo.Echo {
	e := echo.New()
	if container.GetConfig().Logging.Router {
		e.Use(slogecho.New(container.GetLogger()))
	}
	e.Use(middleware.Recover())
	e.Renderer = service.InitTemplates()
	e.Static("/css", "tmpl/css")
	e.Static("/images", "tmpl/images")
	createLoginHandlers(e, container)
	createTemplateHandlers(e, container)
	return e
}

func createTemplateHandlers(e *echo.Echo, container container.Container) {
	RegisterPaths(e, container)
}

func createLoginHandlers(e *echo.Echo, container container.Container) {
	loginHandler := handlers.NewLoginHandler(container)
	e.GET("/auth/callback", loginHandler.Callback)
	e.GET("/auth/login", loginHandler.Login)
	e.POST("/auth/logout", loginHandler.Logout)
}
