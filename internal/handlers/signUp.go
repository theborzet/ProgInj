package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

func (h Handler) SignUpUser(c *fiber.Ctx) error {

	body := models.Client{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Проверяем валидность данных пользователя
	if body.Username == "" || body.Password == "" || body.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username, password and email are required"})
	}

	// Создаем экземпляр структуры User и заполняем его данными из запроса

	// Здесь может быть код для создания экземпляра структуры User из полученных данных

	// Вызываем метод AddClient вашего репозитория
	err := h.repo.AddClient(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Возвращаем статус 200 OK
	return c.SendStatus(fiber.StatusOK)
}
