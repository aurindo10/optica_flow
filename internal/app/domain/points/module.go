package points

import "go.uber.org/fx"

var Module = fx.Provide(
	NewCreatePoints,
	NewFindSellerById,
	NewDeletePoints,
)

