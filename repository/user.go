package repository

import (
	"github.com/happsie/gohtmx/container"
	"github.com/happsie/gohtmx/model"
)

type UserRepository interface {
	Save(user model.TwitchUser, auth model.Auth) error
	Get(ID string) (model.TwitchUser, error)
}

type userRepository struct {
	container container.Container
}

func NewUserRepository(container container.Container) UserRepository {
	return &userRepository{
		container: container,
	}
}

func (ur userRepository) Save(user model.TwitchUser, auth model.Auth) error {
	db := ur.container.GetDB()
	tx := db.MustBegin()
	_, err := tx.NamedExec(`INSERT INTO users 
				(login, display_name, id, profile_image_url, email) 
				VALUES (:login, :display_name, :id, :profile_image_url, :email) 
				ON DUPLICATE KEY UPDATE profile_image_url = :profile_image_url, login = :login, display_name = :display_name`, user)
	_, err = tx.NamedExec(`INSERT INTO user_tokens (user_id, access_token, refresh_token)
	 							VALUES (:user_id, :access_token, :refresh_token)
								ON DUPLICATE KEY UPDATE access_token = :access_token, refresh_token = :refresh_token`, auth)
	err = tx.Commit()
	if err != nil {
		return err
	}
	ur.container.GetLogger().Info("user saved", "user", user)
	return nil
}

func (ur userRepository) Get(ID string) (model.TwitchUser, error) {
	db := ur.container.GetDB()
	user := model.TwitchUser{}
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		return model.TwitchUser{}, err
	}
	return user, nil
}
