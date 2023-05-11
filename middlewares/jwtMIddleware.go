package middlewares

import (
	"mini-project/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

func CreateToken(username string, id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["user_id"] = id
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_KEY))
}

var JWTCustom = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_KEY),
	TokenLookup:   "cookie:jwt",
	AuthScheme:    "user",
})
