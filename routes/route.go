package route

import (
	"mini-project/controller"
	m "mini-project/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	m.LoggerMiddleware(e)

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)

	p := e.Group("/user")
	p.Use(m.JWTCustom)
	p.GET("", controller.GetUserController)
	p.PUT("", controller.UpdateUserController)
	p.DELETE("", controller.DeleteUserController)

	i := e.Group("/tasks")
	i.Use(m.JWTCustom)
	i.GET("", controller.GetTasksController)
	i.GET("/:id", controller.GetTaskController)
	i.POST("", controller.CreateTaskController)
	i.POST("/:id", controller.CreateCollabolatorController)
	i.PUT("/:id", controller.UpdateTaskController)
	i.PUT("/:id/status", controller.UpdateTaskStatusController)
	i.DELETE("/:id", controller.DeleteTaskController)
	i.GET("/category", controller.GetCategoryController)
	i.POST("/category", controller.CreateCategoryController)

	return e
}
