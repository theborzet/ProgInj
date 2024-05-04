package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) ViewBookId(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}
	isAuthenticated, ok := sess.Get("isAuthenticated").(bool)
	if !ok {
		isAuthenticated = false
	}
	userID, ok := sess.Get("uID").(int)
	if !ok {
		userID = 0
	}
	access_level, ok := sess.Get("acessLevel").(int)
	if !ok {
		userID = 0
	}
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
		"ClientId":        userID,
		"Book":            book,
		"Author":          author,
		"IsAuthenticated": isAuthenticated,
		"Access_level":    access_level,
	})
}
