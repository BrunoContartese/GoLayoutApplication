package routes

import (
	"github.com/gofiber/fiber/v2"
	"layoutapp/internal/controllers/users"
)

func SetupUsersModuleRoutes(app *fiber.App) {
	app.Get("/users", users.Index)
	app.Get("/users/:id", users.Show)
	app.Put("/users/:id", users.Update)
	app.Post("/users", users.Store)
	app.Delete("/users/:id", users.Destroy)
	app.Post("/users/:id/restore", users.Restore)
}
