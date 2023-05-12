package usecase

import (
	"mini-project/lib/database"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
)

func GetTasks() (resp []payload.GetTask, err error) {
	tasks, err := database.GetTasks()
	if err != nil {
		return []payload.GetTask{}, err
	}
	for _, value := range tasks {
		var collaborators []payload.Collaborator
		for _, val := range value.Collaborators {
			collaborators = append(collaborators, payload.Collaborator{
				ID:       val.UserID,
				Name:     val.User.Name,
				Username: val.User.Username,
			})
		}

		resp = append(resp, payload.GetTask{
			ID:           value.ID,
			Title:        value.Title,
			Description:  value.Description,
			Status:       value.Status,
			Category:     value.Category.Name,
			Collabolator: collaborators,
			Owner:        value.User.Name,
		})
	}
	return
}

func GetTask(id uint) (resp payload.GetTask, err error) {
	value, err := database.GetTask(id)
	if err != nil {
		return payload.GetTask{}, err
	}
	var collaborators []payload.Collaborator
	for _, val := range value.Collaborators {
		collaborators = append(collaborators, payload.Collaborator{
			ID:       val.UserID,
			Name:     val.User.Name,
			Username: val.User.Username,
		})
	}
	resp = payload.GetTask{
		ID:           value.ID,
		Title:        value.Title,
		Description:  value.Description,
		Status:       value.Status,
		Category:     value.Category.Name,
		Owner:        value.User.Name,
		Collabolator: collaborators,
	}
	return
}

func CreateTask(id uint, req *payload.CreateTask) error {

	if err := database.CreateTask(req, id); err != nil {
		return err
	}
	return nil
}

func UpdateTask(req *payload.CreateTask, id uint, userId uint) error {
	task, err := database.GetTask(id)
	if err != nil {
		return err
	}

	if userId != task.UserID {
		for _, value := range task.Collaborators {
			if value.UserID == userId {
				err = nil
				break
			}
			err = echo.NewHTTPError(http.StatusUnauthorized, "you are prohibited")
		}
	}
	if err != nil {
		return err
	}

	if err := database.UpdateTask(req, id); err != nil {
		return err
	}
	return nil
}

func UpdateTaskStatus(id uint) error {
	task, err := database.GetTask(id)
	if err != nil {
		return err
	}
	if task.Status == "completed" {
		if err := database.UpdateUncompletedStatus(id); err != nil {
			return err
		}
	} else {
		if err := database.UpdateCompletedStatus(id); err != nil {
			return err
		}
	}
	return nil
}

func DeleteTask(id uint, userId uint) error {
	task, err := database.GetTask(id)
	if err != nil {
		return err
	}
	if task.UserID != userId {
		return echo.NewHTTPError(http.StatusUnauthorized, "your are prohibited")
	}
	if err := database.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
