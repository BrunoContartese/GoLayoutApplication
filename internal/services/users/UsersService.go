package users

import (
	"github.com/gofiber/fiber/v2"
	"layoutapp/internal/models"
)

func Index(c *fiber.Ctx) []models.User {
	var users []models.User

	users = append(users, models.User{
		ID:        "asdfioasdf-sdafasdf-123",
		Username:  "bcontartese",
		Password:  "estoesunpassword",
		FirstName: "Bruno",
		LastName:  "Contartese",
		CreatedAt: "2022-04-17",
		UpdatedAt: "2022-04-17",
		DeletedAt: "",
	})

	return users
}

func Show(c *fiber.Ctx) (string, error) {
	return "Busco el usuario con id " + string(c.Params("id")) + " y lo muestro.", nil
}

func Store(c *fiber.Ctx) (string, error) {
	return "Almaceno un usuario.", nil
}

func Update(c *fiber.Ctx) (string, error) {
	return "Busco el usuario con id " + string(c.Params("id")) + " y lo actualizo.", nil
}

func Destroy(c *fiber.Ctx) (string, error) {
	return "Busco el usuario con id " + string(c.Params("id")) + " y lo doy de baja.", nil
}

func Restore(c *fiber.Ctx) (string, error) {
	return "Busco el usuario con id " + string(c.Params("id")) + " y lo doy de alta.", nil
}
