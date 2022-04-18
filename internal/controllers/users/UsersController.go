package users

import (
	"github.com/gofiber/fiber/v2"
)
import "layoutapp/internal/services/users"

func Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(users.Index(c))
}

func Show(c *fiber.Ctx) error {
	res, err := users.Show(c)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON(res)
}

func Store(c *fiber.Ctx) error {
	res, err := users.Store(c)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON(res)
}

func Update(c *fiber.Ctx) error {
	res, err := users.Update(c)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON(res)
}

func Destroy(c *fiber.Ctx) error {
	res, err := users.Destroy(c)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON(res)
}

func Restore(c *fiber.Ctx) error {
	res, err := users.Restore(c)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON(res)
}
