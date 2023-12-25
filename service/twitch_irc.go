package service

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/happsie/gohtmx/container"
)

type TwitchIRC struct {
	container container.Container
	client    *twitch.Client
}

func NewTwitchIRC(container container.Container) *TwitchIRC {
	return &TwitchIRC{
		container: container,
	}
}

func (ti *TwitchIRC) StartIRC(channels []string) error {
	if ti.client != nil || !ti.container.GetConfig().TwitchChatBot.Enabled {
		return nil
	}
	twitchConfig := ti.container.GetConfig().TwitchChatBot
	ti.client = twitch.NewClient(twitchConfig.Username, fmt.Sprintf("oauth:%s", twitchConfig.Token))
	ti.client.OnConnect(ti.onConnect)
	ti.client.OnPrivateMessage(ti.onPrivateMessage)
	ti.client.Join(channels...)
	err := ti.client.Connect()
	if err != nil {
		ti.container.GetLogger().Error("error connecting to IRC", err)
		return err
	}
	ti.container.GetLogger().Info("Connected to IRC")
	return nil
}

func (ti *TwitchIRC) Join(channel string) {
	ti.client.Join(channel)
}

func (ti *TwitchIRC) onConnect() {
	ti.container.GetLogger().Info("connected!")
}

func (ti *TwitchIRC) onPrivateMessage(message twitch.PrivateMessage) {
	ti.container.GetLogger().Info(message.Message)
}
