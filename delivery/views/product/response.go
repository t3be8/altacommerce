package product

import (
	"net/http"

	"github.com/t3be8/altacommerce/entity"
)

func SuccessInsert(data entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data product",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data product",
		"status":  false,
		"data":    nil,
	}
}

func SuccessDelete(data entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil delete product",
		"status":  true,
		"data":    data,
	}
}

func SuccessSelect(data []entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil select product",
		"status":  true,
		"data":    data,
	}
}

func SuccessUpdate(data entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil update stock product",
		"status":  true,
		"data":    data,
	}
}
