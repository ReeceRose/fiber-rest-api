package book

import (
	"github.com/gofiber/fiber/v2"
)

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
