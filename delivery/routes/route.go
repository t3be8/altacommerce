package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/t3be8/altacommerce/delivery/controllers/cart"
	"github.com/t3be8/altacommerce/delivery/controllers/order"
	"github.com/t3be8/altacommerce/delivery/controllers/product"
	"github.com/t3be8/altacommerce/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc user.IUserController, pc product.IProductController, cc cart.ICartController, oc order.IOrderController) {
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
	apiGroup.POST("/products", pc.InsertProduct(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	apiGroup.GET("/products", pc.SelectProduct())
	apiGroup.GET("/products/{id}", pc.GetAllProductById())
	apiGroup.GET("/categories/{categoryId}/products", pc.GetAllProductByCategory())
	apiGroup.PUT("/products/{id}", pc.UpdateProduct())
	apiGroup.DELETE("/products/{id}", pc.DeletedProduct())

	apiGroup.POST("/cart", cc.InsertCart())
	apiGroup.GET("/cart/{id}", cc.SelectCart())
	apiGroup.PUT("/cart/{id}", cc.UpdateCart())
	apiGroup.DELETE("/cart/{id}", cc.DeletedCart())

	apiGroup.POST("/orders", oc.CreateOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	apiGroup.POST("/orders/{order_id}/cancel", oc.CancelOrder(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))
	apiGroup.POST("/orders/{order_id}/payout", oc.Payment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RU$SI4")}))

}
