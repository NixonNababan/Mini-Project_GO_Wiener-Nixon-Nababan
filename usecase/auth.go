package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"

	"golang.org/x/crypto/bcrypt"
)

func Login(req *payload.Login) (models.User, error) {
	user, err := database.GetUserByUsername(req.Username)
	if err != nil {
		return models.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Register(req *payload.Register) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(password),
	}
	if err := database.Register(&user); err != nil {
		return err
	}
	return nil
}
