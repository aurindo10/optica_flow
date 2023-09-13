package controllers

import (
	"errors"
	tradeproduct "optica_flow/internal/app/domain/trade_product"

	"github.com/gofiber/fiber/v2"
)

type TradeProductController struct {
	createTradeProduct *tradeproduct.CreateTradeProduct
	findAllTradeProducts *tradeproduct.FindAllRepository
	updateTradeProduct *tradeproduct.UpdateTradeProduct
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
func (r *TradeProductController) FindAllTradeProducts(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id é obrigatório",
		})
	}
	result, error := r.findAllTradeProducts.Execute(id)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"error interno ao tentar obter os produ†os de troca"})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (r *TradeProductController) UpdateTradeProduct(c *fiber.Ctx) error {
	var request tradeproduct.TradeProducToUpdate
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	result, error := r.updateTradeProduct.Execute(&request)
	if  error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func NewTradeProductController(createTradeProduct *tradeproduct.CreateTradeProduct,
	updateTradeProduct *tradeproduct.UpdateTradeProduct,
	findAllTradeProducts *tradeproduct.FindAllRepository) *TradeProductController {
	return &TradeProductController{
		createTradeProduct: createTradeProduct,
		findAllTradeProducts: findAllTradeProducts,
		updateTradeProduct: updateTradeProduct,
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
