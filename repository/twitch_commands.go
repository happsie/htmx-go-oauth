package repository

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/model"
)

type TwitchCommandRepository interface {
	List(userID string) ([]model.TwitchCommand, error)
}

type twitchCommandRepository struct {
	container container.Container
}

func NewTwitchCommandRepository(container container.Container) TwitchCommandRepository {
	return &twitchCommandRepository{
		container: container,
	}
}

func (ur *twitchCommandRepository) List(userID string) ([]model.TwitchCommand, error) {
	db := ur.container.GetDB()
	commands := []model.TwitchCommand{}
	err := db.Select(&commands, "SELECT * FROM twitch_commands WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return commands, nil
}
