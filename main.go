package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	userController "github.com/t3be8/altacommerce/delivery/controllers/user"
	"github.com/t3be8/altacommerce/delivery/routes"
	"github.com/t3be8/altacommerce/entity"
	userRepo "github.com/t3be8/altacommerce/repository/user"
)

func main() {
	// setup configuration
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(entity.Address{})
	db.AutoMigrate(entity.User{})
	e := echo.New()

	repoUser := userRepo.New(db)

	controllerUser := userController.New(repoUser, validator.New())
	routes.RegisterPath(e, controllerUser)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
