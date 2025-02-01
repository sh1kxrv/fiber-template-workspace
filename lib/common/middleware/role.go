package middleware

import (
	"common/enum"
	"common/errors"
	"common/utils"
	"common/utils/helper"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleMiddleware(role enum.Role) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(utils.JwtClaims)
		if !ok || user.Role != role {
			return helper.SendError(c, nil, errors.Forbidden)
		}
		return c.Next()
	}
}
