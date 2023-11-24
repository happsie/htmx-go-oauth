package model

import "time"

type TwitchCommand struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	Command   string
	Response  string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
