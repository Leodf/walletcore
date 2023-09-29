package router

import (
	"github.com.br/Leodf/walletcore/balances/factory"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/balances/:account_id", factory.MakeBalancesControler)
}
