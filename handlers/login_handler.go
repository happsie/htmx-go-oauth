package handlers

import (
	"fmt"
	"github.com/happise/pixelwars/container"
	"github.com/happise/pixelwars/model"
	"github.com/happise/pixelwars/repository"
	"github.com/happise/pixelwars/service"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
	"slices"
	"time"
)

type LoginHandler interface {
	Login(c echo.Context) error
	Callback(c echo.Context) error
}
type loginHandler struct {
	container      container.Container
	discordService service.TwitchAuthService
	userRepository repository.UserRepository
	jwtService     service.JWTService
	twitchApi      service.TwitchApiService
}

// TODO: append when a new login occurs, delete on successful callback. Make it threadsafe
var states []string

func NewLoginHandler(container container.Container) LoginHandler {
	return &loginHandler{
		container:      container,
		discordService: service.NewTwitchAuthService(container),
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
	token, err := lh.discordService.VerifyCallback(code)
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
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwt
	cookie.Expires = time.Now().Add(72 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/")
	//return c.Render(200, "home.html", nil)
}
