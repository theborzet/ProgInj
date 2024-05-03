package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/library_project/internal/db/models"
)

// var store = session.New()

// func GenerateCSRFToken() string {
// 	token := make([]byte, 32)
// 	_, err := rand.Read(token)
// 	if err != nil {
// 		panic(err) // Обработайте ошибку по вашему усмотрению
// 	}
// 	return base64.URLEncoding.EncodeToString(token)
// }

func (h Handler) SignUpUser(c *fiber.Ctx) error {

	var user models.Client
	user.Username = c.FormValue("username")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")
	if err := h.repo.UserExists(user.Username); err == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "User allready exist"})
	}

	// Проверяем соответствие паролей
	if user.Password != c.FormValue("confirm_password") {
		// Если пароли не совпадают, возвращаем ошибку
		return c.Render("sign_up", fiber.Map{
			"Error": "Пароли не совпадают",
		})
	}

	// Вызываем метод репозитория для добавления пользователя в базу данных
	if err2 := h.repo.AddClient(&user); err2 != nil {
		// Обработка ошибки добавления пользователя
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Возвращаем успешный ответ
	return c.SendStatus(http.StatusOK)
}

// func (h Handler) SignUpUserForm(c *fiber.Ctx) error {
// 	// Создание сессии и генерация токена CSRF
// 	sess, err := store.Get(c)
// 	if err != nil {
// 		return err
// 	}

// 	csrfToken := GenerateCSRFToken()

// 	// Сохранение токена CSRF в сессии
// 	sess.Set("csrf_token", csrfToken)
// 	if err := sess.Save(); err != nil {
// 		return err
// 	}

// 	// Передача токена CSRF в шаблон
// 	return c.Render("sign_up", fiber.Map{
// 		"CSRFToken": csrfToken,
// 	})
// }

// package handlers

// import (
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/csrf"
// 	"github.com/theborzet/library_project/internal/db/models"
// )

// // var store = session.New()

// // func GenerateCSRFToken() string {
// // 	token := make([]byte, 32)
// // 	_, err := rand.Read(token)
// // 	if err != nil {
// // 		panic(err) // Обработайте ошибку по вашему усмотрению
// // 	}
// // 	return base64.URLEncoding.EncodeToString(token)
// // }

// func (h Handler) SignUpUserForm(c *fiber.Ctx) error {
// 	// // Получаем CSRF-токен для текущего запроса

// 	csrf_token := csrf.Token(c)

// 	// // Передаем токен CSRF в шаблон
// 	return c.Render("sign_up", fiber.Map{
// 		"CSRFToken": csrf_token,
// 	})
// }

// func (h Handler) SignUpUser(c *fiber.Ctx) error {
// 	// // Получение сессии
// 	// sess, err := store.Get(c)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// // Получение токена CSRF из сессии
// 	// csrfToken := sess.Get("csrf_token")
// 	// if csrfToken == nil || csrfToken == "" {
// 	// 	// Если токен CSRF отсутствует в сессии или его значение пусто, вернем ошибку "Отсутствие токена CSRF в сессии"
// 	// 	return fiber.NewError(fiber.StatusForbidden, "Отсутствие токена CSRF в сессии")
// 	// }

// 	// // Проверка токена CSRF из формы
// 	// if csrfToken != c.FormValue("csrf_token") {
// 	// 	// Если токены не совпадают, возвращаем ошибку
// 	// 	return fiber.NewError(fiber.StatusForbidden, "Недействительный токен CSRF")
// 	// }

// 	csrfToken := c.Get("X-CSRF-Token")

// 	// Проверяем токен CSRF
// 	// if !csrfToken.Valid() {
// 	// 	// Если токен недействителен, возвращаем ошибку
// 	// 	return fiber.NewError(http.StatusForbidden, "CSRF token is invalid")
// 	// }

// 	var user models.Client

// 	// Парсим данные из тела запроса в структуру UserRegistrationData
// 	if err := c.BodyParser(&user); err != nil {
// 		return err
// 	}

// 	// Проверяем соответствие паролей
// 	if user.Password != c.FormValue("confirm_password") {
// 		// Если пароли не совпадают, возвращаем ошибку
// 		return c.Render("sign_up", fiber.Map{
// 			"CSRFToken": csrfToken,
// 			"Error":     "Пароли не совпадают",
// 		})
// 	}

// 	// Вызываем метод репозитория для добавления пользователя в базу данных
// 	err2 := h.repo.AddClient(&user)
// 	if err2 != nil {
// 		// Обработка ошибки добавления пользователя
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
// 	}

// 	// Возвращаем успешный ответ
// 	return c.SendStatus(http.StatusOK)
// }
