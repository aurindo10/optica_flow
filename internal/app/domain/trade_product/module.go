package tradeproduct

import "go.uber.org/fx"

var Module = fx.Provide(
	NewTradeProduct,
	NewCreateCreateTradeProduct,
	NewFindAllRepository,
	NewUpdateTradeProduct,
)