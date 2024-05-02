package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) SignInUser(c *fiber.Ctx) error {
	var user *models.Client
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.repo.GetPass(username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if password == user.Password {
		return c.Redirect("/index", fiber.StatusFound)
	}
	// return fiber.NewError(fiber.StatusForbidden, "Неверный пароль")
	return c.Render("sign_in", fiber.Map{})
}
