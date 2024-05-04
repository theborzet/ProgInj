package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignInUser(c *fiber.Ctx) error {
	// Получаем сессию из контекста
	sess, err := h.session.Get(c)
	if err != nil {
		return err
	}

	// Получаем пользователя из базы данных
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := h.repo.GetUser(username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Проверяем правильность пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		// Устанавливаем значение "isAuthenticated" в сессии
		sess.Set("isAuthenticated", true)
		sess.Set("uID", user.ID)
		sess.Set("acessLevel", user.AccessLevel)
		// Сохраняем сессию
		if err := sess.Save(); err != nil {
			return err
		}
		// Перенаправляем на главную страницу после успешной аутентификации
		return c.Redirect("/", fiber.StatusFound)
	}

	// В случае неправильного пароля возвращаем шаблон для входа
	return c.Render("sign_in", fiber.Map{})
}
