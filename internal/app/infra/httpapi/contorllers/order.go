package controllers

import (
	"optica_flow/internal/app/domain/orders"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	createOrder *orders.CreateOrder
}

func (r *OrderController) CreateOrder(c *fiber.Ctx)  error {
	body := &orders.Params{}
	error := c.BodyParser(body)
	if error != nil {
		return error
	}
	error = r.ValidateFields(body)
	if error != nil {
		return error
	}
	result, error := r.createOrder.Execute(body)
	if error != nil {
		return error
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (r *OrderController) ValidateFields(result * orders.Params) error {
	if result.CompanyID == "" {
		return fiber.NewError(400, "company_id is required")
	}
	if result.Fase == "" {
		return fiber.NewError(400, "fase is required")
	}
	if result.OrderDate.IsZero() {
		return fiber.NewError(400, "order_date is required")
	}
	if result.ProductName == "" {
		return fiber.NewError(400, "product_name is required")
	}
	if result.Quantity == 0 {
		return fiber.NewError(400, "quantity is required")
	}
	if result.WhoCreatedID == "" {
		return fiber.NewError(400, "who_created_id is required")
	}
	if result.WhoUpdatedID == "" {
		return fiber.NewError(400, "who_updated_id is required")
	}
	return nil
}

func NewOrderController(createOrder *orders.CreateOrder) *OrderController {
	return &OrderController{
		createOrder: createOrder,
	}
}