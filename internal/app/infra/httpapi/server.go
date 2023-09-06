package httpapi

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)


func NewServer(
	lc fx.Lifecycle,
	router *fiber.App,
) *fasthttp.Server {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go router.Listen(":3000")  // Use a instância router aqui
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return router.Shutdown()
		},
	})
	return router.Server()
}
