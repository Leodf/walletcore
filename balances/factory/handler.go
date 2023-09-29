package factory

import (
	"github.com.br/Leodf/walletcore/balances/config"
	database "github.com.br/Leodf/walletcore/balances/database/repository"
	"github.com.br/Leodf/walletcore/balances/domain/usecase"
	"github.com.br/Leodf/walletcore/balances/handler"
	"github.com/gofiber/fiber/v2"
)

func MakeBalancesControler(c *fiber.Ctx) error {
	repository := database.New(config.DB)
	usecase := usecase.NewGetBalancesByAccountId(repository)
	controller := handler.NewBalanceHandler(usecase)
	return controller.GetBalanceHandler(c)
}
