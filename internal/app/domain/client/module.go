package client

import "go.uber.org/fx"

var Module = fx.Provide(
	NewClient,
	NewCreate,
	NewFindClients,
	NewUpdateClient,
	NewFindOne,
	NewDeleteClient,
)