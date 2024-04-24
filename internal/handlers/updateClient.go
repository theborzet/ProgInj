package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var updatedClient models.Client
	if err := c.BodyParser(&updatedClient); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	client := models.Client{
		Username:    updatedClient.Username,
		Password:    updatedClient.Password,
		Email:       updatedClient.Email,
		AccessLevel: updatedClient.AccessLevel,
		Books:       updatedClient.Books,
	}

	errchan := make(chan error)

	go func() {
		if err := h.repo.UpdateClient(uint(updatedId), &client); err != nil {
			errchan <- err
		} else {
			errchan <- nil
		}

	}()
	err = <-errchan
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&client)
}
