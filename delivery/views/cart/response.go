package cart

import (
	"net/http"

	"github.com/t3be8/altacommerce/entity"
)

func SuccessInsert(data entity.Cart) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data cart",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data cart",
		"status":  false,
		"data":    nil,
	}
}

func SuccessDelete(data entity.Cart) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil delete cart",
		"status":  true,
		"data":    data,
	}
}

func SuccessSelect(data []entity.Cart) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil select cart",
		"status":  true,
		"data":    data,
	}
}

func SuccessUpdate(data entity.Cart) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil update stock cart",
		"status":  true,
		"data":    data,
	}
}
