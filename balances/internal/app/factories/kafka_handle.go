package factories

import (
	"github.com.br/Leodf/walletcore/balances/internal/domain/usecase"
	"github.com.br/Leodf/walletcore/balances/internal/infra/database/connection"
	database "github.com.br/Leodf/walletcore/balances/internal/infra/database/mysql"
)

func MakeKafkaHandle() *usecase.CreateBalances {
	repo := database.New(connection.DB)
	return usecase.NewCreateBalances(repo)
}
