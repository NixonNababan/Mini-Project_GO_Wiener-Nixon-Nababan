package controller

import (
	"mini-project/models/payload"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetTasksController(c echo.Context) error {
	data, err := usecase.GetTasks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all tasks",
		"data":    data,
	})
}

func GetTaskController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := usecase.GetTask(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get task",
		"data":    data,
	})
}

func CreateTaskController(c echo.Context) error {
	req := payload.CreateTask{}
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := GetUserId(c)
	err := usecase.CreateTask(id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create task",
	})
}

func UpdateTaskController(c echo.Context) error {
	var req payload.CreateTask
	id, _ := strconv.Atoi(c.Param("id"))
	userId := GetUserId(c)
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateTask(&req, uint(id), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update task",
	})
}

func UpdateTaskStatusController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := usecase.UpdateTaskStatus(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update status task",
	})
}

func DeleteTaskController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userId := GetUserId(c)
	err := usecase.DeleteTask(uint(id), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete task",
	})
}
