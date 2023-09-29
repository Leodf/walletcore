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

// type repositoryInput struct {
// 	ID           string
// 	AccountID    string
// 	Amount       string
// 	LastUpdate   time.Time
// 	ID_2         string
// 	AccountID_2  string
// 	Amount_2     string
// 	LastUpdate_2 time.Time
// }

type BalanceRepository interface {
	CheckAccountIdExists(ctx context.Context, accountID string) (bool, error)
	// SaveBalances(ctx context.Context, balances repositoryInput) error
	FindBalancesByAccountId(ctx context.Context, accountId string) ([]RepositoryOutput, error)
}
