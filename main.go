package main

import (
	"fmt"

	// "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/config"
)

func main() {
	// setup configuration
	conf := config.InitConfig()
	db := config.InitDB(*conf)

	defer db.Close()

	var version string
	res := db.QueryRow("select version()").Scan(&version)
	if res != nil {
		log.Fatal(res)
	}
	fmt.Println(version)

	defer res.Close()
	// e := echo.New()
	// log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
