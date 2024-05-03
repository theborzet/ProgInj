package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jmoiron/sqlx"
	"github.com/theborzet/library_project/internal/db/repository"
)

type Handler struct {
	repo     repository.Repository
	sessions *session.Session
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h Handler) IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func RegistrationRoutess(app *fiber.App, db *sqlx.DB) {
	handler := NewHandler(repository.NewSQLRepository(db))

	app.Get("/", handler.IndexView)

	bookRoutes := app.Group("/books")
	bookRoutes.Post("/add", handler.AddBook)
	bookRoutes.Get("/add", func(c *fiber.Ctx) error {
		authors, err := handler.repo.GetAllAuthors("", "")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Render("add_book", fiber.Map{
			"Authors": authors,
			"Title":   "Добавление книги",
		})
	})
	bookRoutes.Delete("/:id", handler.DeleteBook)
	bookRoutes.Get("/", handler.ViewAllBooks)
	bookRoutes.Get("/:id", handler.ViewBookId)
	bookRoutes.Put("/:id", handler.UpdateBook)
	bookRoutes.Post("/", handler.ViewAllBooks)

	authorRoutes := app.Group("/authors")
	authorRoutes.Post("/add", handler.AddAuthor)
	authorRoutes.Get("/add", func(c *fiber.Ctx) error {
		return c.Render("add_author", fiber.Map{
			"Title": "Добавление автора",
		})
	})
	authorRoutes.Delete("/:id", handler.DeleteAuthor)
	authorRoutes.Get("/", handler.ViewAllAuthors)
	authorRoutes.Get("/:id", handler.ViewAuthorId)
	authorRoutes.Put("/:id", handler.UpdateAuthor)
	authorRoutes.Post("/", handler.ViewAllAuthors)

	clientRoutes := app.Group("/client")
	clientRoutes.Post("/", handler.AddClient)
	clientRoutes.Delete("/:id", handler.DeleteClient)
	clientRoutes.Get("/", handler.ViewAllClients)
	clientRoutes.Get("/:id", handler.ViewClientId)
	clientRoutes.Put("/:id", handler.UpdateClient)

	app.Get("/registration", func(c *fiber.Ctx) error {
		return c.Render("sign_up", fiber.Map{})
	})
	app.Post("/registration", handler.SignUpUser)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("sign_in", fiber.Map{})
	})
	app.Post("/login", handler.SignInUser)

}
