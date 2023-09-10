package controllers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewProductController,
	NewFornecedorController,
	NewClientController,
	NewOrderController,
	NewProductOrderController,
)