package tradeproductrepository

import (
	tradeproduct "optica_flow/internal/app/domain/trade_product"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewTradeProductRepository,
		fx.As(new(tradeproduct.Repository)),
	),
)