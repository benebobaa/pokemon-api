package app

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"pokemon-api/util"
)

func NewFiber(c util.Config, controller *Controller) *fiber.App {
	config := fiber.Config{
		AppName:           c.AppName,
		EnablePrintRoutes: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
	}

	if c.GOEnv == "production" {
		config.Prefork = true
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Static("/public/images", "./public/images")

	NewRouter(app, controller)

	return app
}
