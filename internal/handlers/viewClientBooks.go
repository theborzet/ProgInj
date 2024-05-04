package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
	"github.com/theborzet/library_project/internal/pagination"
)

func (h Handler) ViewClientBooks(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}
	userID, ok := sess.Get("uID").(int)
	if !ok {
		userID = 0
	}
	isAuthenticated, ok := sess.Get("isAuthenticated").(bool)
	if !ok {
		isAuthenticated = false
	}
	var books []*models.Book

	books, err = h.repo.ViewClientBook(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	client, err := h.repo.GetClientID(uint(userID))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	pageSize := 3
	paginatedBooks, paginator := pagination.PaginateBooks(books, page, pageSize)

	return c.Render("profile", fiber.Map{
		"Client":          client,
		"Books":           paginatedBooks,
		"Paginator":       paginator,
		"Title":           "Список книг",
		"IsPaginated":     paginator.TotalPages > 1,
		"IsAuthenticated": isAuthenticated,
	})
}
