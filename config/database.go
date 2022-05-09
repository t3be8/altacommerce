package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
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
