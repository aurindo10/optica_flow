package routers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewProductRouter,
	MakeRouter,
	NewFornecedorRouter,
	NewClientRouter,
	NewOrderRouter,
	NewProductOrderRouter,
	NewPointsRouter,
	NewTradeProductController,
	NewComissionRouter,
)