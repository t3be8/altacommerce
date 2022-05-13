package order

import "net/http"

type ResponsePayOrder struct {
	StatusCode  string `json:"status_code"`
	PaymentType string `json:"payment_type"`
	Status      string `json:"status"`
}

type ResponseOrder struct {
	ID       uint    `json:"id"`
	Address  string  `json:"address"`
	TotalPay float64 `json:"total_pay"`
	Status   string  `json:"status"`
}

func StatusDetailOrder(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Order",
		"status":  true,
		"data":    data,
	}
}

func StatusCreated(OrderID string, snap map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Order",
		"status":  true,
		"data":    map[string]interface{}{"order-id": OrderID, "RedirectUrl": snap},
	}
}

func StatusPayOrder(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Pay Order",
		"status":  true,
		"data":    data,
	}
}

func StatusCancelOrder() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Cancel Order",
		"status":  true,
	}
}

func StatusSnapError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNoContent,
		"message": "Error Get Urls",
		"status":  false,
	}
}

func StatusUpdateOrder(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Update Order Status",
		"status":  true,
		"data":    data,
	}
}
