package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	cartController "github.com/t3be8/altacommerce/delivery/controllers/cart"
	userController "github.com/t3be8/altacommerce/delivery/controllers/user"
	"github.com/t3be8/altacommerce/delivery/routes"
	"github.com/t3be8/altacommerce/entity"
	cartRepo "github.com/t3be8/altacommerce/repository/cart"
	userRepo "github.com/t3be8/altacommerce/repository/user"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(entity.Address{})
	db.AutoMigrate(entity.User{})
	db.AutoMigrate(entity.Cart{})
	e := echo.New()

	repoUser := userRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	routes.RegisterPath(e, controllerUser)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

	db.AutoMigrate(entity.Cart{})

	repocart := cartRepo.New(db)

	controllerCart := cartController.New(repocart, validator.New())
	routes.CartPath(e, controllerCart)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
