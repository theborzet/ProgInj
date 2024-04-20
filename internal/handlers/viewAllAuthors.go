package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) ViewAllAuthors(c *fiber.Ctx) error {
	authors, err := h.repo.GetAllAuthors()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(authors)
}
