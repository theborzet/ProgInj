package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) ViewAllBooks(c *fiber.Ctx) error {
	books, err := h.repo.GetAllBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(books)
}
