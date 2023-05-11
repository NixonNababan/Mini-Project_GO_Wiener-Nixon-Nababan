package controller

import (
	"mini-project/lib/cookie"
	"mini-project/middlewares"
	"mini-project/models/payload"
	"mini-project/usecase"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	var login payload.Login
	c.Bind(&login)

	if err := c.Validate(login); err != nil {
		return err
	}

	user, err := usecase.Login(&login)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := middlewares.CreateToken(user.Username, user.ID)
	if err != nil {
		return err
	}

	cookie.CreateJWTCookies(c, token, "jwt")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
	})
}

func RegisterController(c echo.Context) error {
	var registerForm payload.Register
	c.Bind(&registerForm)

	if err := c.Validate(registerForm); err != nil {
		return err
	}

	err := usecase.Register(&registerForm)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
	})
}

func GetUserId(c echo.Context) uint {
	cookie, _ := c.Cookie("jwt")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	id := uint(claims["user_id"].(float64))
	return id
}
