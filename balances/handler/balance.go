package handler

import (
	"github.com.br/Leodf/walletcore/balances/domain/exception"
	"github.com.br/Leodf/walletcore/balances/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type BalanceHandler struct {
	getBalanceByAccountId *usecase.GetBalancesByAccountId
}

func NewBalanceHandler(gb *usecase.GetBalancesByAccountId) *BalanceHandler {
	return &BalanceHandler{
		getBalanceByAccountId: gb,
	}
}

func (uc *BalanceHandler) GetBalanceHandler(c *fiber.Ctx) error {
	accountId := c.Params("account_id")
	inputDto := usecase.GetBalancesInput{
		AccountID: accountId,
	}
	balances, err := uc.getBalanceByAccountId.Execute(inputDto)
	if err == exception.ErrUserNotExits {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": exception.ErrInternalServerError.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": true,
		"data":   balances,
		"err":    nil,
	})
}
