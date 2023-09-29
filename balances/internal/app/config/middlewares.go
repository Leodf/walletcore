package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMidlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
}
