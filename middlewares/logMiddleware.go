package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "status = ${status}, time = ${time_rfc3339}, method = ${method}, uri = ${ur}, \n",
	}))
}
