package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetPdfBook(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	userID, ok := sess.Get("uID").(int)
	if !ok {
		userID = 0
	}
	id := c.Params("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	h.repo.AddClientBook(userID, bookId)

	book, err := h.repo.GetBookID(uint(bookId))
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	c.Redirect("/static/vendor/pdf/" + book.Title + ".pdf")

	return nil
}
