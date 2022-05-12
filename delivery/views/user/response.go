package user

import (
	"net/http"
)

type LoginResponse struct {
	Data  UserResponse
	Token string
}

type UserResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Phone *string `json:"phone"`
}

func RegisterSuccess(data UserResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil register user baru",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data user",
		"status":  false,
		"data":    nil,
	}
}
