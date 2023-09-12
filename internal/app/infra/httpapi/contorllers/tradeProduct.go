package controllers

import (
	"errors"
	tradeproduct "optica_flow/internal/app/domain/trade_product"

	"github.com/gofiber/fiber/v2"
)

type TradeProductController struct {
	createTradeProduct *tradeproduct.CreateTradeProduct
}

func (f *TradeProductController) CreateTradeProduct(c *fiber.Ctx) error {
	var request tradeproduct.TradeProductParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Houve algum error em objter o body"})
	}
	if error := f.Validate(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	result, error := f.createTradeProduct.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error ao criar produto de troca"})
	}
	return c.Status(fiber.StatusOK).JSON(result)
} 

func NewTradeProductController(createTradeProduct *tradeproduct.CreateTradeProduct) *TradeProductController {
	return &TradeProductController{
		createTradeProduct: createTradeProduct,
	}
}

func (c *TradeProductController) Validate(p *tradeproduct.TradeProductParams) error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Description == "" {
		return errors.New("description is required")
	}
	if p.CompanyID == "" {
		return errors.New("companyID is required")
	}
	if p.PointAmmount == 0 {
		return errors.New("pointAmmount is required")
	}
	if p.WhoCreatedID == "" {
		return errors.New("whoCreatedID is required")
	}
	return nil
}
