package comissionvalue

import "go.uber.org/fx"


var Module = fx.Provide(
	NewComissionValues,
	NewCreateValueComission,
	NewFindAllByCompanyId,
	NewDeleteComissionVAlue,
)