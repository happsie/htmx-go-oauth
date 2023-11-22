package handlers

import (
	"fmt"
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/model"
	"github.com/happsie/gohtmx/repository"
	"github.com/happsie/gohtmx/service"
	"github.com/happsie/gohtmx/utils"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
	"slices"
)

type LoginHandler interface {
	Login(c echo.Context) error
	Callback(c echo.Context) error
	Logout(c echo.Context) error
}
type loginHandler struct {
	container      container.Container
	twitchService  service.AuthService
	userRepository repository.UserRepository
	jwtService     service.JWTService
	twitchApi      service.TwitchApiService
}

// TODO: append when a new login occurs, delete on successful callback. Make it threadsafe
var states []string

func NewLoginHandler(container container.Container) LoginHandler {
	return &loginHandler{
		container:      container,
		twitchService:  service.NewAuthService(container),
		userRepository: repository.NewUserRepository(container),
		jwtService:     service.NewJWTService(container),
		twitchApi:      service.NewTwitchApiService(container),
	}
}

func (lh loginHandler) Login(c echo.Context) error {
	state, err := gonanoid.New()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	states = append(states, state)
	c.Response().Header().Add("HX-Redirect", lh.container.GetOauthConfig().AuthCodeURL(state))
	return c.NoContent(302)
}

func (lh loginHandler) Callback(c echo.Context) error {
	if !slices.Contains(states, c.FormValue("state")) {
		lh.container.GetLogger().Error("invalid state", "url", c.Request().URL.Path)
		return c.NoContent(http.StatusUnauthorized)
	}
	code := c.FormValue("code")
	token, err := lh.twitchService.VerifyCallback(code)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	user, err := lh.twitchApi.FetchCurrentUser(token.AccessToken)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	auth := model.Auth{
		UserID:       user.ID,
		AccessToken:  fmt.Sprintf("%s %s", token.TokenType, token.AccessToken),
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	err = lh.userRepository.Save(user, auth)
	if err != nil {
		lh.container.GetLogger().Error("could not save user", "error", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}
	jwt, err := lh.jwtService.CreateJWT(user, token.Expiry.Unix())
	if err != nil {
		lh.container.GetLogger().Error("error creating JWT", "error", err)
		return c.NoContent(http.StatusUnauthorized)
	}
	utils.SetAuthCookie(c, jwt)
	return c.Redirect(http.StatusSeeOther, "/home")
}

func (lh loginHandler) Logout(c echo.Context) error {
	utils.InvalidateAuthCookie(c)
	c.Response().Header().Add("HX-Redirect", "/")
	return c.NoContent(302)
}
