package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var LIMIT_ACCOUNT = -500

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
	if b.Amount < float64(LIMIT_ACCOUNT) {
		return errors.New("amount exceeds limit")
	}
	return nil
}
