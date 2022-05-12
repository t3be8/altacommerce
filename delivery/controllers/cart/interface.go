package cart

import "github.com/labstack/echo/v4"

type ICartController interface {
	InsertCart() echo.HandlerFunc
	SelectCart() echo.HandlerFunc
	DeletedCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
}
