package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config AppConfig) *gorm.DB {

	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
		&entity.Address{}, &entity.Shipment{}, &entity.OrderDetail{}, &entity.Order{}, &entity.Category{}, &entity.Product{}, &entity.Cart{})

}
