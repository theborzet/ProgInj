package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	clientID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	errchan := make(chan error)
	go func() {
		if err := h.repo.DeleteRecord("client", clientID); err != nil {
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
