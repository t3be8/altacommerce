package view

import "net/http"

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "terdapat kesalahan pada server",
		"status":  false,
		"data":    nil,
	}
}

func OK(data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
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
