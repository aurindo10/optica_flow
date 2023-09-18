package cashflowout

import "go.uber.org/fx"



var Module = fx.Provide(
	NewCashFlowBalance,
	NewCreateFlowBalance,
	NewFindByRangeDate,
	NewUpdateCashFlowBalance,
	NewDeleteCashFlowBalance,
)