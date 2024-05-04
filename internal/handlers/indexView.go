package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) IndexView(c *fiber.Ctx) error {
	// Получаем сессию из контекста
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}
	userID, ok := sess.Get("uID").(int)
	if !ok {
		userID = 0
	}

	// Получаем значение "isAuthenticated" из сессии
	isAuthenticated, ok := sess.Get("isAuthenticated").(bool)
	if !ok {
		isAuthenticated = false
	}

	// Рендеринг шаблона index.tmpl с передачей значения isAuthenticated
	return c.Render("index", fiber.Map{
		"isAuthenticated": isAuthenticated,
		"ClientId":        userID,
	})
}
