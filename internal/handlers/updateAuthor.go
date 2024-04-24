package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var updatedAuthor models.Author
	if err := c.BodyParser(&updatedAuthor); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	author := models.Author{
		FirstName: updatedAuthor.FirstName,
		LastName:  updatedAuthor.LastName,
		BirthDate: updatedAuthor.BirthDate,
	}

	errchan := make(chan error)

	go func() {
		if err := h.repo.UpdateAuthor(uint(updatedId), &author); err != nil {
			errchan <- err
		} else {
			errchan <- nil
		}

	}()
	err = <-errchan
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&author)
}
