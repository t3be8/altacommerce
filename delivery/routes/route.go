package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t3be8/altacommerce/delivery/controllers/cart"
	"github.com/t3be8/altacommerce/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc user.IUserController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/login", uc.Login())
	e.POST("/register", uc.Register())

}

func CartPath(e *echo.Echo, cc cart.ICartController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/cart", cc.InsertCart())
	e.GET("/cart/{id}", cc.SelectCart())
	e.PUT("/cart/{id}", cc.UpdateCart())
	e.DELETE("/cart/{id}", cc.DeletedCart())
}
