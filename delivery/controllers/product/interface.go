package product

import "github.com/labstack/echo/v4"

type IProductController interface {
	InsertProduct() echo.HandlerFunc
	SelectProduct() echo.HandlerFunc
	DeletedProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
}
