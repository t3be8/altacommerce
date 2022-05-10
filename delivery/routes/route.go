package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t3be8/altacommerce/delivery/controllers/product"
	"github.com/t3be8/altacommerce/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc user.IUserController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	apiGroup := e.Group("/api/v1")

	apiGroup.POST("/login", uc.Login())
	apiGroup.POST("/register", uc.Register())
}

func ProductPath(e *echo.Echo, pc product.IProductController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/product", pc.InsertProduct())
	e.GET("/product", pc.SelectProduct())
	e.GET("/product/{id}", pc.GetAllProductById())
	e.GET("/categories/{categoryId}/product", pc.GetAllProductByCategory())
	e.PUT("/product/{id}", pc.UpdateProduct())
	e.DELETE("/product/{id}", pc.DeletedProduct())
}
