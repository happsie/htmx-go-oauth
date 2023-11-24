package service

import (
	"context"
	"github.com/happsie/gohtmx/container"
	"golang.org/x/oauth2"
)

type AuthService interface {
	VerifyCallback(code string) (*oauth2.Token, error)
}

type authService struct {
	container container.Container
}

func NewAuthService(container container.Container) AuthService {
	return &authService{
		container: container,
	}
}

func (ds authService) VerifyCallback(code string) (*oauth2.Token, error) {
	oauth := ds.container.GetOauthConfig()
	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		ds.container.GetLogger().Error(err.Error())
		return nil, err
	}
	return token, nil
}
