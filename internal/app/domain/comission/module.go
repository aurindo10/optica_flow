package comission

import "go.uber.org/fx"

var Module = fx.Provide(
	NewComission,
	NewCreateComission,
	NewFindByUserId,
	NewDeleteComission,
)