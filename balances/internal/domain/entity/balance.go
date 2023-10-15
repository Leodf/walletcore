package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var MIN_LIMIT_ACCOUNT = 5

type Balance struct {
	ID         string
	AccountID  string
	Amount     float64
	LastUpdate time.Time
}

func NewBalance(accountId string, amount float64) (*Balance, error) {
	balance := &Balance{
		ID:         uuid.NewString(),
		AccountID:  accountId,
		Amount:     amount,
		LastUpdate: time.Now(),
	}
	err := balance.validate()
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *Balance) validate() error {
	if b.AccountID == "" {
		return errors.New("accountId cannot be empty")
	}
	if b.Amount < float64(MIN_LIMIT_ACCOUNT) {
		return errors.New("amount must be greater than 5, but the core service saves an amount in the account table until it reaches a zero value - this is a simulated error")
	}
	return nil
}
