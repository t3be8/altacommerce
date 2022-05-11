package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	productController "github.com/t3be8/altacommerce/delivery/controllers/product"
	userController "github.com/t3be8/altacommerce/delivery/controllers/user"
	"github.com/t3be8/altacommerce/delivery/routes"
	"github.com/t3be8/altacommerce/entity"
	productRepo "github.com/t3be8/altacommerce/repository/product"
	userRepo "github.com/t3be8/altacommerce/repository/user"
)

func main() {
	// setup configuration
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(&entity.User{}, &entity.Address{}, &entity.Category{}, &entity.Product{})

	e := echo.New()

	repoUser := userRepo.New(db)
	repoProduct := productRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	controllerProduct := productController.New(repoProduct, validator.New())

	routes.RegisterPath(e, controllerUser, controllerProduct)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
