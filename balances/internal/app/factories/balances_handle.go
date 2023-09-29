package factories

import (
	handler "github.com.br/Leodf/walletcore/balances/internal/app/handlers"
	"github.com.br/Leodf/walletcore/balances/internal/database/connection"
	database "github.com.br/Leodf/walletcore/balances/internal/database/mysql"
	"github.com.br/Leodf/walletcore/balances/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

func MakeBalanceHandler() func(*fiber.Ctx) error {
	repository := database.New(connection.Db)
	usecase := usecase.NewGetBalancesByAccountId(repository)
	controller := handler.NewBalanceHandler(usecase)
	return controller.GetBalanceHandler
}
