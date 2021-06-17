package book

import (
	"github.com/ReeceRose/fiber-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)

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

	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.SendStatus(500)
		return c.SendString("No book found with given ID")
	}

	db.Delete(&book)
	return c.JSON("Book successfully deleted")
}
