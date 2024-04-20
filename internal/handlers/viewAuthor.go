package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) ViewAuthorId(c *fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	var author *models.Author

	author, err = h.repo.GetAuthorID(authorId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&author)

}
