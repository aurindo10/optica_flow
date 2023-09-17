package flowbalancerepository

import (
	cashflowout "optica_flow/internal/app/domain/cash_flow_out"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewFlowBalanceRepository,
		fx.As(new(cashflowout.Repository)),
	),
)