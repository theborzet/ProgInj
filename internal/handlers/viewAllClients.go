package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) ViewAllClients(c *fiber.Ctx) error {
	clients, err := h.repo.GetAllClients()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(clients)
}
