package routes

import (
	"github.com/gofiber/fiber/v2"
	"layoutapp/internal/controllers/users"
)

func SetupUsersModuleRoutes(app *fiber.App) {
	app.Get("/users", users.Index)
}
