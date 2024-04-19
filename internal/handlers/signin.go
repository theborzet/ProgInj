package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) SignInUser(c *fiber.Ctx) error {
	var user models.Client
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Проверяем валидность данных пользователя
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username, password and email are required"})
	}

	// Создаем экземпляр структуры User и заполняем его данными из запроса

	// Здесь может быть код для создания экземпляра структуры User из полученных данных

	// Вызываем метод AddClient вашего репозитория
	err := h.repo.AddClient(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Возвращаем статус 200 OK
	return c.SendStatus(fiber.StatusOK)
}
