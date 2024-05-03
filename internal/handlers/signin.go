package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SignInUser(c *fiber.Ctx) error {
	// Получаем сессию из контекста
	// sess, err := h.sessions.Get(c)
	// if err != nil {
	// 	return err
	// }

	// Получаем пользователя из базы данных
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := h.repo.GetPass(username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Проверяем правильность пароля
	if password == user.Password {
		// Устанавливаем значение "userID" в сессии
		// sess.Set("userID", user.ID)

		// Сохраняем сессию
		// if err := sess.Save(); err != nil {
		// 	return err
		// }

		// Перенаправляем на главную страницу после успешной аутентификации
		return c.Redirect("/", fiber.StatusFound)
	}

	// В случае неправильного пароля возвращаем шаблон для входа
	return c.Render("sign_in", fiber.Map{})
}
