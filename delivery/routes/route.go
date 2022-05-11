package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t3be8/altacommerce/delivery/controllers/product"
	"github.com/t3be8/altacommerce/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc user.IUserController, pc product.IProductController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	apiGroup := e.Group("/api/v1")

	// user login and register enpoints route
	apiGroup.POST("/login", uc.Login())
	apiGroup.POST("/register", uc.Register())

	// product enpoints route
	apiGroup.POST("/product", pc.InsertProduct())
	apiGroup.GET("/product", pc.SelectProduct())
	apiGroup.GET("/product/{id}", pc.GetAllProductById())
	apiGroup.GET("/categories/{categoryId}/product", pc.GetAllProductByCategory())
	apiGroup.PUT("/product/{id}", pc.UpdateProduct())
	apiGroup.DELETE("/product/{id}", pc.DeletedProduct())
}
