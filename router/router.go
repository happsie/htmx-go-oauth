package router

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/handlers"
	"github.com/happsie/gohtmx/repository"
	"github.com/happsie/gohtmx/service"
	"github.com/happsie/gohtmx/utils"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
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
	userRepo := repository.NewUserRepository(container)
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
	commandHandler := handlers.NewCommandHandler(container)
	e.GET("/commands", commandHandler.List)
}

func createLoginHandlers(e *echo.Echo, container container.Container) {
	loginHandler := handlers.NewLoginHandler(container)
	e.GET("/auth/callback", loginHandler.Callback)
	e.GET("/auth/login", loginHandler.Login)
	e.POST("/auth/logout", loginHandler.Logout)
}
