package controllers

import (
	"errors"
	cashflowout "optica_flow/internal/app/domain/cash_flow_out"

	"github.com/gofiber/fiber/v2"
)


type CashFlowBalanceController struct {
	CreateCashFlowBalance *cashflowout.CreateFlowBalance
}

func (r *CashFlowBalanceController) CreateFlowBalance(c *fiber.Ctx)  error {
	var request cashflowout.CashFlowBalanceParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	if error := r.Validate(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	result, error := r.CreateCashFlowBalance.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).
		JSON(error.Error())
	}
	return c.Status(fiber.StatusCreated).
	JSON(result)
}

func NewCashFlowBalanceController(CreateCashFlowBalance *cashflowout.CreateFlowBalance) *CashFlowBalanceController {
	return &CashFlowBalanceController{
		CreateCashFlowBalance: CreateCashFlowBalance,
	}
}

func (r *CashFlowBalanceController) Validate(p *cashflowout.CashFlowBalanceParams) error {
	if p.CompanyID == "" {
		return errors.New("o campo company Id é obrigatório")
	}
	if p.Description == "" {
		return errors.New("o campo description  é obrigatório")
	}
	if p.Type == "" {
		return errors.New("o campo type  é obrigatório")
	}
	if p.WhoCreatedID == "" {
		return errors.New("o campo whocreatedid  é obrigatório")
	}
	if p.WhoUpdatedID == "" {
		return errors.New("o campo whoupdatedid  é obrigatório")
	}
	if p.Value == 0.0 {
		return errors.New("o campo value Id é obrigatório")
	}
	if p.DueDate.IsZero() {
		return errors.New("o campo duedate Id é obrigatório")
	}
	return nil
}