package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	book := models.Book{
		Title:           updatedBook.Title,
		AuthorID:        updatedBook.AuthorID,
		PublicationYear: updatedBook.PublicationYear,
		Genre:           updatedBook.Genre,
		Count:           updatedBook.Count,
		ImageUrl:        updatedBook.ImageUrl,
	}

	errchan := make(chan error)

	go func() {
		if err := h.repo.UpdateBook(updatedId, &book); err != nil {
			errchan <- err
		} else {
			errchan <- nil
		}

	}()
	err = <-errchan
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&book)
}
