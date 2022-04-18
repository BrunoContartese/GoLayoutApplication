package users

import (
	"github.com/gofiber/fiber/v2"
)
import "layoutapp/internal/services/users"

func Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(users.Index())
}
