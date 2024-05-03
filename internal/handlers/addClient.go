package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) AddClient(c *fiber.Ctx) error {
	body := models.Client{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	client := models.Client{
		Username:    body.Username,
		Password:    body.Password,
		Email:       body.Email,
		AccessLevel: body.AccessLevel,
	}

	errchan := make(chan error)
	go func() {
		err := h.repo.AddClient(&client)
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

	return c.Status(fiber.StatusCreated).JSON(&client)
}
