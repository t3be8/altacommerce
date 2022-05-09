package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t3be8/altacommerce/delivery/controllers/product"
)

func ProductPath(e *echo.Echo, pc product.IProductController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/product", pc.InsertProduct())
	e.GET("/product", pc.SelectProduct())
	e.PUT("/product/{id}", pc.UpdateProduct())
	e.DELETE("/product/{id}", pc.DeletedProduct())
}
