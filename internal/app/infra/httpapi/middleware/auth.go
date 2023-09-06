package middleware

import (
	env "optica_flow/pkg"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gofiber/fiber/v2"
)
type AuthMiddleware struct {
	clerkSecret string
}
func (p *AuthMiddleware) Authenticate(c *fiber.Ctx) error {
	client, err := clerk.NewClient(p.clerkSecret)
	if err != nil {
		return fiber.NewError(500, "Error generating clerk client")
	}
	sessionToken := c.Get("Authorization")
	if sessionToken == "" {
		return fiber.NewError(500, "Error getting token")
	}
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")
	_, err = client.VerifyToken(sessionToken)
	if err != nil {
		return fiber.NewError(500, "Error verifying token")
	}
	return c.Next()
}
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		clerkSecret: env.Config("CLERK_SECRET_KEY"),
	}
}