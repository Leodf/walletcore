package config

import (
	"github.com.br/Leodf/walletcore/balances/internal/app/factories"
	"github.com.br/Leodf/walletcore/balances/internal/infra/http/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello world!")
	})
	api := app.Group("/api")
	routes.NewBalanceRoutes(api, *factories.MakeBalanceHandle())
}
