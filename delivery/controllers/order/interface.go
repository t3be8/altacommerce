package order

import "github.com/labstack/echo/v4"

type IOrderController interface {
	CreateOrder() echo.HandlerFunc
	CancelOrder() echo.HandlerFunc
	Payment() echo.HandlerFunc
}
