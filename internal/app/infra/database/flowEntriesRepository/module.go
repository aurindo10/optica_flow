package flowentriesrepository

import (
	flowentries "optica_flow/internal/app/domain/flow_entries"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewFlowEntrieRepository,
		fx.As(new(flowentries.Repository)),
	),
)