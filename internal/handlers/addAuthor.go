package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) AddAuthor(c *fiber.Ctx) error {
	body := models.Author{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	author := models.Author{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		BirthDate: body.BirthDate,
	}

	errchan := make(chan error)
	go func() {
		err := h.repo.AddAuthor(&author)
		if err != nil {
			errchan <- fiber.NewError(fiber.StatusBadRequest, err.Error())
		} else {
			errchan <- nil
		}
	}()

	err := <-errchan
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&author)
}
