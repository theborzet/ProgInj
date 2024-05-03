package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) LogoutUser(c *fiber.Ctx) error {
	// Получаем сессию из контекста
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}

	// Удаляем значения сесии
	sess.Delete("isAuthenticated")
	sess.Delete("uID")
	sess.Delete("acessLevel")

	// Сохраняем сессию
	if err := sess.Save(); err != nil {
		return err
	}

	// Перенаправление на главную страницу
	return c.Redirect("/", fiber.StatusFound)
}
