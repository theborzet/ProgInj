package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignInUser(c *fiber.Ctx) error {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	return c.SendStatus(fiber.StatusOK)
}
