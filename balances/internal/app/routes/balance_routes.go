package routes

import (
	"github.com.br/Leodf/walletcore/balances/internal/app/factories"
	"github.com/gofiber/fiber/v2"
)

func BalanceRoutes(route fiber.Router) {
	route.Get("/balances/:account_id", factories.MakeBalanceHandler())
}
