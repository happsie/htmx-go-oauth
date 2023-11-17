package service

import (
	"encoding/json"
	"fmt"
	"github.com/happise/pixelwars/container"
	"github.com/happise/pixelwars/model"
	"io"
	"net/http"
)

type TwitchApiService interface {
	FetchCurrentUser(accessToken string) (model.TwitchUser, error)
}

type twitchApi struct {
	container container.Container
}

func NewTwitchApiService(container container.Container) TwitchApiService {
	return &twitchApi{
		container: container,
	}
}

func (ta twitchApi) FetchCurrentUser(accessToken string) (model.TwitchUser, error) {
	clientID := ta.container.GetOauthConfig().ClientID
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	if err != nil {
		ta.container.GetLogger().Error("could not create request", "error", err.Error())
		return model.TwitchUser{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Client-Id", clientID)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ta.container.GetLogger().Error("could not fetch user", "error", err.Error())
		return model.TwitchUser{}, err
	}
	if res.StatusCode != 200 {
		ta.container.GetLogger().Error("could not fetch user", "status", res.StatusCode)
		return model.TwitchUser{}, fmt.Errorf("could not fetch user")
	}
	user, err := ta.convertBodyToUser(res)
	return user, err
}

func (ta twitchApi) convertBodyToUser(r *http.Response) (model.TwitchUser, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ta.container.GetLogger().Error("could not read body", "error", err.Error())
		return model.TwitchUser{}, err
	}
	defer r.Body.Close()
	type DataWrapper struct {
		Data []model.TwitchUser `json:"data"`
	}
	var dw DataWrapper
	err = json.Unmarshal(body, &dw)
	if err != nil {
		ta.container.GetLogger().Error("could not unmarshal json", "error", err.Error())
		return model.TwitchUser{}, err
	}
	if len(dw.Data) != 1 {
		ta.container.GetLogger().Error("could not find user")
		return model.TwitchUser{}, fmt.Errorf("could not find user")
	}
	return dw.Data[0], nil
}
