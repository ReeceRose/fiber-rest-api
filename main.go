package main

import (
	"github.com/ReeceRose/fiber-rest-api/book"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
