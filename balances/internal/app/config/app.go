package config

import (
	"log"

	"github.com.br/Leodf/walletcore/balances/pkg/env"
	"github.com/gofiber/fiber/v2"
)

func InitApp() {
	app := fiber.New()
	SetupMidlewares(app)
	SetupRoutes(app)
	err := app.Listen(env.GetEnv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
