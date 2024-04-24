package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) ViewBookId(c *fiber.Ctx) error {
	id := c.Params("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	var book *models.Book

	book, err = h.repo.GetBookID(uint(bookId))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	author, err := h.repo.GetAuthorID(book.AuthorID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Render("book_detail", fiber.Map{
		"Book":   book,
		"Author": author,
	})
}
