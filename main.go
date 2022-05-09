package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
	ProductController "github.com/t3be8/altacommerce/delivery/controllers/product"
	"github.com/t3be8/altacommerce/delivery/routes"
	"github.com/t3be8/altacommerce/entity"
	productRepo "github.com/t3be8/altacommerce/repository/product"
)

func main() {
	// setup configuration
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(entity.Product{})
	db.AutoMigrate(entity.ProductCategory{})
	e := echo.New()

	repoProduct := productRepo.New(db)

	controllerProduct := ProductController.New(repoProduct, validator.New())
	routes.ProductPath(e, controllerProduct)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
