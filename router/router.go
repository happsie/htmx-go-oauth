package router

import (
	"github.com/happise/pixelwars/container"
	"github.com/happise/pixelwars/handlers"
	"github.com/happise/pixelwars/repository"
	"github.com/happise/pixelwars/service"
	"github.com/happise/pixelwars/utils"

	//echojwt "github.com/labstack/echo-jwt/v4"
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
	//e.Use(echojwt.WithConfig(GetJwtConfig(container))) Set this on routes instead! https://github.com/labstack/echo/issues/1737
	e.Static("/css", "tmpl/css")
	e.Static("/images", "tmpl/images")
	createLoginHandlers(e, container)
	createTemplateHandlers(e, container)
	return e
}

func createTemplateHandlers(e *echo.Echo, container container.Container) {
	userRepo := repository.NewUserRepository(container)
	e.GET("/", func(c echo.Context) error {
		_, err := c.Cookie("token")
		auth, err := utils.GetAuthInfo(c, container.GetConfig())
		// TODO: Also check if expired
		if err != nil {
			container.GetLogger().Error("could not authenticate", err.Error())
			return c.Render(200, "index.html", nil)
		}
		user, err := userRepo.Get(auth.UserId)
		if err != nil {
			return c.Render(200, "index.html", nil)
		}
		return c.Render(200, "home.html", user)
	})
}

func createLoginHandlers(e *echo.Echo, container container.Container) {
	loginHandler := handlers.NewLoginHandler(container)
	e.GET("/api/auth/callback", loginHandler.Callback)
	e.GET("/api/auth/login", loginHandler.Login)
}
