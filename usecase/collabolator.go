package usecase

import "mini-project/lib/database"

func CreateCollabolator(taskId uint, userId uint) error {
	if err := database.CreateCollabolator(taskId, userId); err != nil {
		return err
	}
	return nil
}
