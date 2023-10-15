package usecase

import (
	"context"

	"github.com.br/Leodf/walletcore/balances/internal/domain/entity"
	"github.com.br/Leodf/walletcore/balances/internal/domain/repository"
)

type CreateBalancesInput struct {
	AccountID   string
	Amount      float64
	AccountID_2 string
	Amount_2    float64
}

type CreateBalances struct {
	repository repository.BalanceRepository
}

func NewCreateBalances(br repository.BalanceRepository) *CreateBalances {
	return &CreateBalances{
		repository: br,
	}
}

func (uc *CreateBalances) Execute(input CreateBalancesInput) error {
	ctx := context.Background()
	balanceFrom, err := entity.NewBalance(input.AccountID, input.Amount)
	if err != nil {
		return err
	}
	balanceTo, err := entity.NewBalance(input.AccountID_2, input.Amount_2)
	if err != nil {
		return err
	}
	balances := repository.RepositoryInput{
		ID:           balanceFrom.ID,
		AccountID:    balanceFrom.AccountID,
		Amount:       balanceFrom.Amount,
		LastUpdate:   balanceFrom.LastUpdate,
		ID_2:         balanceTo.ID,
		AccountID_2:  balanceTo.AccountID,
		Amount_2:     balanceTo.Amount,
		LastUpdate_2: balanceTo.LastUpdate,
	}
	err = uc.repository.SaveBalances(ctx, balances)
	if err != nil {
		return err
	}
	return nil
}
