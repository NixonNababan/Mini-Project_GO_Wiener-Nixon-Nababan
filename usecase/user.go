package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(id uint) (payload.GetUser, error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return payload.GetUser{}, err
	}
	resp := payload.GetUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	return resp, nil
}

func UpdateUser(id uint, req *payload.UpdateUser) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(password),
	}
	if err := database.UpdateUser(&user, id); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := database.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
