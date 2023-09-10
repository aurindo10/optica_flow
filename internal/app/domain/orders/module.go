package orders

import "go.uber.org/fx"

var Module = fx.Provide(
	NewOrder,
	NewCreateOrder,
	NewFindOne,
	NewUpdateOrder,
	NewFindAllOrder,
	NewDeleteOrder,
)

