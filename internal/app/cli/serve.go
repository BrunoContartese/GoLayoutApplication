package cli

import (
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"salescrm/internal/routes"
)

func ServeApplication() {
	var port string
	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	/*Load .env file*/
	loadEnviromentFile()

	/*Load middlewares*/
	loadMiddlewares(app)

	/*Initialize cors*/
	initializeCors(app)

	/*Setup routes*/
	setupRoutes(app)

	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	color.Green("Server is running on port: " + port)

	log.Fatal(app.Listen(":" + port))
}

func loadEnviromentFile() {
	color.Green("Loading enviroment file...")
	// load .env file
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

/*Application middlewares*/
func loadMiddlewares(app *fiber.App) {
	color.Green("Loading middlewares...")
}

/*Initialize cors*/
func initializeCors(app *fiber.App) {
	color.Green("Initializing cors...")
	app.Use(cors.New())
}

/*Setup routes*/
func setupRoutes(app *fiber.App) {
	color.Green("Setting up routes...")
	routes.SetupRoutes(app)
}
