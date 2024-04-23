package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) AddBook(c *fiber.Ctx) error {
	body := models.Book{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	book := models.Book{
		Title:           body.Title,
		AuthorID:        body.AuthorID,
		PublicationYear: body.PublicationYear,
		Genre:           body.Genre,
		Count:           body.Count,
		ImageUrl:        body.ImageUrl,
	}

	errchan := make(chan error)
	go func() {
		err := h.repo.AddBook(&book)
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

	return c.Status(fiber.StatusCreated).JSON(&book)
}
