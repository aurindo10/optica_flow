package controllers

import (
	productorder "optica_flow/internal/app/domain/product_order"

	"github.com/gofiber/fiber/v2"
)

type ProductOrderController struct {
	createProductOrder *productorder.CreateProductOrder 
}

func (r *ProductOrderController) Create(c *fiber.Ctx) error {
	p := &productorder.ProductOrderParms{}
	if err := c.BodyParser(p); err != nil {
		return err
	}
	error := r.Validate(p)
	if error != nil {
		return error
	}
	result, error := r.createProductOrder.Execute(p)
	if error != nil {
		return error
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (r *ProductOrderController) Validate(c *productorder.ProductOrderParms) error {
	if c.Amout == 0 {
		return fiber.NewError(400, "Quantity is required")
	}
	if c.OrderID == nil {
		return fiber.NewError(400, "Order is required")
	}
	if c.ProductID == nil {
		return fiber.NewError(400, "Product is required")
	}
	return nil
}
func NewProductOrderController(createProductOrder *productorder.CreateProductOrder) *ProductOrderController {
	return &ProductOrderController{createProductOrder: createProductOrder}
}