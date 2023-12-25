package handlers

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/repository"
	"github.com/happsie/gohtmx/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommandHandler interface {
	List(c echo.Context) error
	Modal(c echo.Context) error
}

type commandHandler struct {
	container         container.Container
	twitchCommandRepo repository.TwitchCommandRepository
	userRepo          repository.UserRepository
}

func NewCommandHandler(container container.Container) CommandHandler {
	return &commandHandler{
		container:         container,
		twitchCommandRepo: repository.NewTwitchCommandRepository(container),
		userRepo:          repository.NewUserRepository(container),
	}
}

func (ch *commandHandler) List(c echo.Context) error {
	auth, err := utils.GetAuthInfo(c, ch.container.GetConfig())
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	user, err := ch.userRepo.Get(auth.UserId)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	commands, err := ch.twitchCommandRepo.List(user.ID)
	if err != nil {
		ch.container.GetLogger().Error("could not load commands", "error", err)
		return c.Render(200, "commands-error", nil)
	}
	return c.Render(200, "commands", commands)
}

func (ch *commandHandler) Modal(c echo.Context) error {
	return c.Render(200, "command-modal", nil)
}
