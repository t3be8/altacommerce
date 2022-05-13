package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	cartController "github.com/t3be8/altacommerce/delivery/controllers/cart"
	orderController "github.com/t3be8/altacommerce/delivery/controllers/order"
	productController "github.com/t3be8/altacommerce/delivery/controllers/product"
	userController "github.com/t3be8/altacommerce/delivery/controllers/user"
	"github.com/t3be8/altacommerce/delivery/routes"
	"github.com/t3be8/altacommerce/entity"
	cartRepo "github.com/t3be8/altacommerce/repository/cart"
	orderRepo "github.com/t3be8/altacommerce/repository/order"
	productRepo "github.com/t3be8/altacommerce/repository/product"
	userRepo "github.com/t3be8/altacommerce/repository/user"
	"github.com/t3be8/altacommerce/utils"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(
		&entity.User{},
		&entity.Address{}, &entity.Shipment{}, &entity.OrderDetail{}, &entity.Order{}, &entity.Category{}, &entity.Product{}, &entity.Cart{})

	snap := utils.InitMidrans()

	e := echo.New()

	repoUser := userRepo.New(db)
	repoProduct := productRepo.New(db)
	repocart := cartRepo.New(db)
	repoOrder := orderRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	controllerOrder := orderController.New(repoOrder, snap, validator.New())
	controllerProduct := productController.New(repoProduct, validator.New())
	controllerCart := cartController.New(repocart, validator.New())
	routes.RegisterPath(e, controllerUser, controllerProduct, controllerCart, controllerOrder)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
