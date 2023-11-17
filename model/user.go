package model

import "time"

type TwitchUser struct {
	Login           string
	DisplayName     string `db:"display_name"`
	ProfileImageUrl string `db:"profile_image_url"`
	Email           string
	ID              string
}

type Auth struct {
	UserID       string `db:"user_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	Expiry       time.Time
}

type JwtInfo struct {
	UserId   string
	Username string
	Exp      int64
}
