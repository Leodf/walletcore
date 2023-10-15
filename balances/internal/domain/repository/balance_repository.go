package repository

import (
	"context"
	"time"
)

type RepositoryOutput struct {
	ID         string
	AccountID  string
	Amount     float64
	LastUpdate time.Time
}

type RepositoryInput struct {
	ID           string
	AccountID    string
	Amount       float64
	LastUpdate   time.Time
	ID_2         string
	AccountID_2  string
	Amount_2     float64
	LastUpdate_2 time.Time
}

type BalanceRepository interface {
	CheckAccountIdExists(ctx context.Context, accountID string) (bool, error)
	SaveBalances(ctx context.Context, balances RepositoryInput) error
	FindBalancesByAccountId(ctx context.Context, accountId string) ([]RepositoryOutput, error)
}
