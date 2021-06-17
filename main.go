package main

import (
	"fmt"

	"github.com/ReeceRose/fiber-rest-api/controllers"
	"github.com/ReeceRose/fiber-rest-api/database"
	"github.com/ReeceRose/fiber-rest-api/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", controllers.GetBooks)
	app.Get("/api/v1/book/:id", controllers.GetBook)
	app.Post("/api/v1/book", controllers.NewBook)
	app.Delete("/api/v1/book/:id", controllers.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Datbase connection successfully opened")

	// Auto migrate DB
	database.DBConn.AutoMigrate(&entities.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
