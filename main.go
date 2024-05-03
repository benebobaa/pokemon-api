package main

import (
	"github.com/gofiber/fiber/v2/log"
	"pokemon-api/app"
	"pokemon-api/util"
)

func main() {
	viperConfig, err := util.LoadConfig(".")
	db := app.NewDatabaseConnection(viperConfig.DBDsn)
	validate := app.NewValidator()

	controller := app.NewController(db, validate)
	fiber := app.NewFiber(viperConfig, controller)

	err = fiber.Listen(":" + viperConfig.PortApp)
	if err != nil {
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
