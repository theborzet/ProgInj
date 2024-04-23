package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/pagination"
)

func (h Handler) ViewAllBooks(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	books, err := h.repo.GetAllBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pageSize := 10
	paginatedBooks, paginator := pagination.PaginateBooks(books, page, pageSize)

	return c.Render("book_list", fiber.Map{
		"Books":       paginatedBooks,
		"Paginator":   paginator,
		"Title":       "Список книг",
		"IsPaginated": paginator.TotalPages > 1,
	})
}
