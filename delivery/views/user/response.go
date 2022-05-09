package user

import (
	"net/http"

	"github.com/t3be8/altacommerce/entity"
)

type LoginResponse struct {
	Data  entity.User
	Token string
}

func SuccessInsert(data entity.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data pegawai",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data pegawai",
		"status":  false,
		"data":    nil,
	}
}
