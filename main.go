package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	cartController "github.com/t3be8/altacommerce/delivery/controllers/cart"
  productController "github.com/t3be8/altacommerce/delivery/controllers/product"
	userController "github.com/t3be8/altacommerce/delivery/controllers/user"
	"github.com/t3be8/altacommerce/delivery/routes"
	productRepo "github.com/t3be8/altacommerce/repository/product"
	userRepo "github.com/t3be8/altacommerce/repository/user"
  cartRepo "github.com/t3be8/altacommerce/repository/cart"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)

	db.AutoMigrate(&entity.User{}, &entity.Address{}, &entity.Category{}, &entity.Product{}, &entity.Cart{})

	e := echo.New()

	repoUser := userRepo.New(db)
	repoProduct := productRepo.New(db)
  repocart := cartRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	controllerProduct := productController.New(repoProduct, validator.New())
  controllerCart := cartController.New(repocart, validator.New())
	routes.RegisterPath(e, controllerUser, controllerProduct , controllerCart)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
