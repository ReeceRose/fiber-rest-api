package controllers

import (
	"github.com/ReeceRose/fiber-rest-api/database"
	"github.com/ReeceRose/fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []models.Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book models.Book
	db.Find(&book, id)

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": err})
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book models.Book
	db.First(&book, id)

	if book.Title == "" {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "No book found with given ID"})
	}

	db.Delete(&book)
	return c.JSON(fiber.Map{"message": "Book successfully deleted"})
}
