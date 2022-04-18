package routes

import (
	"github.com/gofiber/fiber/v2"
	"layoutapp/internal/controllers/users"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", users.Index)
}
