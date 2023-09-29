package config

import (
	"github.com/gofiber/fiber/v2"
)

func InitApp() *fiber.App {
	app := fiber.New()
	SetupMidlewares(app)
	SetupRoutes(app)

	return app
}
