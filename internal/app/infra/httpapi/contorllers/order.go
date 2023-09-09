package controllers

import (
	"optica_flow/internal/app/domain/orders"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderController struct {
	createOrder *orders.CreateOrder
	findOneOrder *orders.FindOne
	updateOrder *orders.UpdateOrder
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
func (r *OrderController) FindOneOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "id is required")
	}
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return fiber.NewError(400, "id is invalid")
	}
	result, error := r.findOneOrder.Execute(idParsed)
	if error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (r *OrderController) UpdateOrder(c *fiber.Ctx) error {
	body := &orders.OrderToUpdate{}
	error := c.BodyParser(body)
	if error != nil {
		return error
	}
	result, error := r.updateOrder.Execute(body)
	if error != nil {
		return error
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func NewOrderController(createOrder *orders.CreateOrder, findOneOrder *orders.FindOne, updateOrder *orders.UpdateOrder) *OrderController {
	return &OrderController{
		createOrder: createOrder,
		findOneOrder: findOneOrder,
		updateOrder: updateOrder,
	}
}