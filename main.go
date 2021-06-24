package main

import (
	"fmt"

	"github.com/ReeceRose/fiber-rest-api/database"
	"github.com/ReeceRose/fiber-rest-api/models"
	"github.com/ReeceRose/fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Datbase connection successfully opened")

	// Auto migrate DB
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	initDatabase()
	defer database.DBConn.Close()

	routes.Setup(app)

	app.Listen(":3000")
}
