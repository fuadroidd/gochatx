package main

import (
	"fmt"
	"gochatx/internal/infrastructure/db"
	"gochatx/internal/infrastructure/db/migrations"

	"gochatx/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	dbcon, err := db.Connect()

	if err != nil {
		fmt.Print(fmt.Errorf("unable to connect database %s", err))
	} else {

		if err := migrations.AutoMigrate(dbcon); err != nil {
			fmt.Print(fmt.Errorf("migrations error %s", err))
		}
	}

	route := gin.Default()

	router.UserRoutes(route, dbcon)

	if err := route.Run("localhost:8000"); err != nil {
		fmt.Print(err.Error())
	}
}
