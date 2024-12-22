package main

import (
	"divrhino-trivia-crud/internal/database"
	"divrhino-trivia-crud/internal/handlers"
	"divrhino-trivia-crud/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {

	database.ConnectDb()

	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/", "./web/public") // Static files

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Use(handlers.NotFound) // 404 page

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
