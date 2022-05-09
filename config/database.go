package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

func InitDB(config AppConfig) *sql.DB {

	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.DBPort,
		config.DBName,
	)

	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatal(err.Error())
	}

	// defer db.Close()

	return db
}
