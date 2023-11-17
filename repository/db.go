package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/happise/pixelwars/config"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const schema = `
CREATE TABLE IF NOT EXISTS users (
	id varchar(255) PRIMARY KEY NOT NULL,
	profile_image_url varchar(255) NOT NULL,
	login varchar(255) NOT NULL,
	display_name varchar(255) NOT NULL,
	email varchar(255) NOT NULL
);
`

/* TODO: Not working for some reason to put in schema.
CREATE TABLE IF NOT EXISTS user_tokens (
user_id VARCHAR(255) PRIMARY KEY NOT NULL,
access_token VARCHAR(100) NOT NULL,
refresh_token VARCHAR(100) NOT NULL,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
*/

func NewDatabase(log *slog.Logger, config config.Config) (*sqlx.DB, error) {
	log.Info("connecting to database")
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@/%s", config.Database.User, config.Database.Password, config.Database.Database))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	log.Info("database connection established")
	return db, nil
}
