package usecase

import (
	"mini-project/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func CreateCollabolator(taskId uint, userId uint) error {
	task, err := database.GetTask(taskId)
	if err != nil {
		return err
	}
	for _, value := range task.Collaborators {
		if value.UserID == userId {
			return echo.NewHTTPError(http.StatusBadRequest, "you have join the task")
		}
	}
	if err := database.CreateCollabolator(taskId, userId); err != nil {
		return err
	}
	return nil
}
