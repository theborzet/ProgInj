package handlers

import "github.com/gofiber/fiber/v2"

func (h Handler) LogoutUser(c *fiber.Ctx) error {
	// Очистка данных сессии
	h.session.Delete("isAuthenticated")
	h.session.Delete("uID")
	h.session.Delete("acessLevel")

	// Перенаправление на главную страницу
	return c.Redirect("/", fiber.StatusFound)
}
