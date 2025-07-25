package routes

import (
	"github.com/gofiber/fiber/v2"

	"book-library/controllers"
	"book-library/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth routes
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// Book routes (public)
	api.Get("/books", controllers.GetAllBooks)
	api.Get("/books/:id", controllers.GetBookByID)

	// Book routes (author only)
	api.Post("/books", middlewares.Protect(), middlewares.AuthorOnly(), controllers.CreateBook)
	api.Put("/books/:id", middlewares.Protect(), middlewares.AuthorOnly(), controllers.UpdateBook)
	api.Delete("/books/:id", middlewares.Protect(), middlewares.AuthorOnly(), controllers.DeleteBook)

	// Reactions (reader only)
	api.Post("/books/:id/react", middlewares.Protect(), middlewares.ReaderOnly(), controllers.ReactToBook)
	api.Delete("/books/:id/react", middlewares.Protect(), middlewares.ReaderOnly(), controllers.RemoveReaction)
}
