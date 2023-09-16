package flowentries

import "go.uber.org/fx"


var Module = fx.Provide(
	NewCashFlowEntries,
	NewCreateFlowEntrie,
	NewFindByIntervalDate,
)