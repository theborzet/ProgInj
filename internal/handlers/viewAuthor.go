package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
	"github.com/theborzet/library_project/internal/pagination"
)

func (h Handler) ViewAuthorId(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	var author *models.Author

	author, err = h.repo.GetAuthorID(uint(authorId))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	books, err := h.repo.GetAuthorBooks(uint(authorId))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pageSize := 10
	paginatedBooks, paginator := pagination.PaginateBooks(books, page, pageSize)

	return c.Render("author_detail", fiber.Map{
		"Author":      author,
		"Books":       paginatedBooks,
		"Paginator":   paginator,
		"IsPaginated": paginator.TotalPages > 1,
	})

}
