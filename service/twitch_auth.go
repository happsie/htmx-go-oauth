package service

import (
	"context"
	"github.com/happise/pixelwars/container"
	"golang.org/x/oauth2"
)

type TwitchAuthService interface {
	VerifyCallback(code string) (*oauth2.Token, error)
}

type twitchAuthService struct {
	container container.Container
}

func NewTwitchAuthService(container container.Container) TwitchAuthService {
	return &twitchAuthService{
		container: container,
	}
}

func (ds twitchAuthService) VerifyCallback(code string) (*oauth2.Token, error) {
	oauth := ds.container.GetOauthConfig()
	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		ds.container.GetLogger().Error(err.Error())
		return nil, err
	}
	return token, nil
}
