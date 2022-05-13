package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId float64, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["email"] = email
	claims["expired"] = time.Now().Add(time.Hour * 2).Unix() //Token expires after 2h
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RU$SI4"))
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}

func ExtractTokenEmailUser(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		return email
	}
	return ""
}
