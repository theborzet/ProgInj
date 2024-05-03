package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

func (h Handler) SignUpUser(c *fiber.Ctx) error {
	// Логируем полученные данные из формы
	log.Printf("Received form data: username=%s, email=%s", c.FormValue("username"), c.FormValue("email"))

	var user models.Client
	user.Username = c.FormValue("username")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")

	// Проверяем, что все обязательные поля формы заполнены
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Println("Missing required fields")
		return c.Render("sign_up", fiber.Map{
			"Error": "All form fields must be filled",
		})
	}

	// Проверяем соответствие паролей
	if user.Password != c.FormValue("confirm_password") {
		// Если пароли не совпадают, возвращаем ошибку
		log.Println("Passwords do not match")
		return c.Render("sign_up", fiber.Map{
			"Error": "Passwords do not match",
		})
	}

	// Хэшируем пароль с использованием bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}
	user.Password = string(hashedPassword)

	if err := h.repo.UserExists(user.Username); err == nil {
		log.Println("User already exists")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "User already exists"})
	}

	// Вызываем метод репозитория для добавления пользователя в базу данных
	if err := h.repo.AddClient(&user); err != nil {
		// Обработка ошибки добавления пользователя
		log.Printf("Failed to register user: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Перенаправляем на страницу авторизации
	log.Println("User registered successfully")
	// Возвращаем успешный ответ
	return c.Redirect("/login", fiber.StatusFound)
}
