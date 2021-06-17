package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Ttile  string  `json:"title"`
	Author string  `json:"author"`
	Rating float32 `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete book")
}
