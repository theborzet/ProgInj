package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/theborzet/library_project/internal/db/repository"
)

type Handler struct {
	repo repository.Repository
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
	bookRoutes.Post("/", handler.AddBook)
	bookRoutes.Delete("/:id", handler.DeleteBook)
	bookRoutes.Get("/", handler.ViewAllBooks)
	bookRoutes.Get("/:id", handler.ViewBookId)
	bookRoutes.Put("/:id", handler.UpdateBook)

	authorRoutes := app.Group("/author")
	authorRoutes.Post("/", handler.AddAuthor)
	authorRoutes.Delete("/:id", handler.DeleteAuthor)
	authorRoutes.Get("/", handler.ViewAllAuthors)
	authorRoutes.Get("/:id", handler.ViewAuthorId)
	authorRoutes.Put("/:id", handler.UpdateAuthor)

	clientRoutes := app.Group("/client")
	clientRoutes.Post("/", handler.AddClient)
	clientRoutes.Delete("/:id", handler.DeleteClient)
	clientRoutes.Get("/", handler.ViewAllClients)
	clientRoutes.Get("/:id", handler.ViewClientId)
	clientRoutes.Put("/:id", handler.UpdateClient)
}
