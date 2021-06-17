package routes

import (
	"github.com/ReeceRose/fiber-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	app.Get("/api/v1/book", controllers.GetBooks)
	app.Get("/api/v1/book/:id", controllers.GetBook)
	app.Post("/api/v1/book", controllers.NewBook)
	app.Delete("/api/v1/book/:id", controllers.DeleteBook)
}
