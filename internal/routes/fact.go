package routes

import (
	"divrhino-trivia-crud/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func FactRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)

	app.Get("/fact/:id/edit", handlers.EditFact)
	app.Patch("/fact/:id", handlers.UpdateFact)

	app.Delete("/fact/:id", handlers.DeleteFact)
}
