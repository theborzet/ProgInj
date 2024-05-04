package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

func (h Handler) SignUpUser(c *fiber.Ctx) error {
	var user models.Client
	user.Username = c.FormValue("username")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")

	// Проверяем, что все обязательные поля формы заполнены
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Render("sign_up", fiber.Map{
			"Error": "All form fields must be filled",
		})
	}

	// Проверяем соответствие паролей
	if user.Password != c.FormValue("confirm_password") {
		// Если пароли не совпадают, возвращаем ошибку
		return c.Render("sign_up", fiber.Map{
			"Error": "Passwords do not match",
		})
	}

	// Хэшируем пароль с использованием bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}
	user.Password = string(hashedPassword)

	if err := h.repo.UserExists(user.Username); err == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "User already exists"})
	}

	// Вызываем метод репозитория для добавления пользователя в базу данных
	if err := h.repo.AddClient(&user); err != nil {
		// Обработка ошибки добавления пользователя
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Возвращаем успешный ответ
	return c.Redirect("/login", fiber.StatusFound)
}
