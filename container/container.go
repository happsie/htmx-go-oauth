package container

import (
	"github.com/happsie/gohtmx/config"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"log/slog"
)

type Container interface {
	GetConfig() config.Config
	GetOauthConfig() *oauth2.Config
	GetLogger() *slog.Logger
	GetDB() *sqlx.DB
}
type container struct {
	config config.Config
	logger *slog.Logger
	db     *sqlx.DB
}

func NewContainer(config config.Config, logger *slog.Logger, db *sqlx.DB) Container {
	return &container{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (c *container) GetConfig() config.Config {
	return c.config
}

func (c *container) GetOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  c.GetConfig().Oauth.RedirectURL,
		ClientID:     c.GetConfig().Oauth.ClientID,
		ClientSecret: c.GetConfig().Oauth.ClientSecret,
		Scopes:       []string{"user:read:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://id.twitch.tv/oauth2/authorize",
			TokenURL: "https://id.twitch.tv/oauth2/token",
		},
	}
}

func (c *container) GetLogger() *slog.Logger {
	return c.logger
}

func (c *container) GetDB() *sqlx.DB {
	return c.db
}
