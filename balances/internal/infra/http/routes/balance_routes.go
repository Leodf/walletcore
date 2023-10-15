package routes

import (
	"github.com.br/Leodf/walletcore/balances/internal/infra/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiber.Router
	balanceHandler handlers.BalanceHandler
}

func NewBalanceRoutes(router fiber.Router, balanceHandler handlers.BalanceHandler) (*Router, error) {
	router.Get("/balances/:account_id", balanceHandler.GetBalanceHandler)
	return &Router{}, nil
}
