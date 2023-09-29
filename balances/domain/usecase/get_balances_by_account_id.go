package usecase

import (
	"context"
	"time"

	"github.com.br/Leodf/walletcore/balances/domain/exception"
	"github.com.br/Leodf/walletcore/balances/domain/repository"
)

type GetBalancesInput struct {
	AccountID string `json:"accountId"`
}

type GetBalancesOutput struct {
	BalanceID  string    `json:"balanceId"`
	Amount     float64   `json:"amount"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type GetBalancesByAccountId struct {
	repository repository.BalanceRepository
}

func NewGetBalancesByAccountId(br repository.BalanceRepository) *GetBalancesByAccountId {
	return &GetBalancesByAccountId{
		repository: br,
	}
}

func (uc *GetBalancesByAccountId) Execute(input GetBalancesInput) ([]*GetBalancesOutput, error) {
	ctx := context.Background()
	exists, err := uc.repository.CheckAccountIdExists(ctx, input.AccountID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, exception.ErrUserNotExits
	}
	var balances []*GetBalancesOutput
	dataBalances, err := uc.repository.FindBalancesByAccountId(ctx, input.AccountID)
	if err != nil {
		return nil, err
	}

	for _, dataBalance := range dataBalances {
		getBalancesItem := &GetBalancesOutput{
			BalanceID:  dataBalance.ID,
			Amount:     dataBalance.Amount,
			LastUpdate: dataBalance.LastUpdate,
		}
		balances = append(balances, getBalancesItem)
	}

	return balances, nil
}
