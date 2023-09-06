package routers

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Load()
}
func MakeRouter(
	peopleRouter *ProductRouter,
) *fiber.App {
	r := fiber.New()
	peopleRouter.Load(r)
	return r
}