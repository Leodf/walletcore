package factories

import (
	"github.com.br/Leodf/walletcore/balances/internal/domain/usecase"
	"github.com.br/Leodf/walletcore/balances/internal/infra/database/connection"
	database "github.com.br/Leodf/walletcore/balances/internal/infra/database/mysql"
	"github.com.br/Leodf/walletcore/balances/internal/infra/http/handlers"
)

func MakeBalanceHandle() *handlers.BalanceHandler {
	repo := database.New(connection.DB)
	usecase := usecase.NewGetBalancesByAccountId(repo)
	handler := handlers.NewBalanceHandler(usecase)
	return handler
}
