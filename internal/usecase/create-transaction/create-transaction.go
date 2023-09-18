package createTransaction

import (
	"context"

	"github.com.br/Leodf/walletcore/internal/entity"
	"github.com.br/Leodf/walletcore/internal/gateway"
	"github.com.br/Leodf/walletcore/pkg/events"
	unitofwork "github.com.br/Leodf/walletcore/pkg/unit-of-work"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"accountIdFrom"`
	AccountIDTo   string  `json:"accountIdTo"`
	Amount        float64 `json:"amount"`
}
type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"accountIdFrom"`
	AccountIDTo   string  `json:"accountIdTo"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                unitofwork.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	uow unitofwork.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	err := uc.Uow.Do(ctx, func(_ *unitofwork.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}
		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount
		return nil
	})
	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
