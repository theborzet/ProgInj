package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	errchan := make(chan error)
	go func() {
		if err := h.repo.DeleteRecord("author", authorID); err != nil {
			errchan <- err
		} else {
			errchan <- nil
		}
	}()

	err = <-errchan
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
