package controllers

import (
	"github.com/ReeceRose/fiber-rest-api/database"
	"github.com/ReeceRose/fiber-rest-api/entities"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []entities.Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book entities.Book
	db.Find(&book, id)

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(entities.Book)

	if err := c.BodyParser(book); err != nil {
		c.SendStatus(500)
		return c.JSON(err)
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book entities.Book
	db.First(&book, id)

	if book.Title == "" {
		c.SendStatus(500)
		return c.SendString("No book found with given ID")
	}

	db.Delete(&book)
	return c.JSON("Book successfully deleted")
}
