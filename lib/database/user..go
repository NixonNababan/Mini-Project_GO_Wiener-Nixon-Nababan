package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetUserByUsername(id string) (user models.User, err error) {
	if err := config.DB.Where("username = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}

func GetUserById(id uint) (user models.User, err error) {
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return
}

func UpdateUser(req *models.User, id uint) error {
	if err := config.DB.Model(&req).Where("id = ?", id).Updates(models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
