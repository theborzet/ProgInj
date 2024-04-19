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

func RegistrationRoutess(app *fiber.App, db *sqlx.DB) {
	handler := NewHandler(repository.NewSQLRepository(db))
	// bookRoutes := app.Group("/books")
	// bookRoutes.Post("/", handler.AddBook)
	// bookRoutes.Delete("/:id", handler.DeleteProduct)
	// bookRoutes.Get("/", handler.GetAllProduct)
	// bookRoutes.Get("/:id", handler.GetProduct)
	// bookRoutes.Put("/:id", handler.UpdateProduct)

	// authorRoutes := app.Group("/author")
	// bookRoutes.Post("/", handler.AddBook)
	// bookRoutes.Delete("/:id", handler.DeleteProduct)
	// bookRoutes.Get("/", handler.GetAllProduct)
	// bookRoutes.Get("/:id", handler.GetProduct)
	// bookRoutes.Put("/:id", handler.UpdateProduct)

	// clientRoutes := app.Group("/client")
	// bookRoutes.Post("/", handler.AddBook)
	// bookRoutes.Delete("/:id", handler.DeleteProduct)
	// bookRoutes.Get("/", handler.GetAllProduct)
	// bookRoutes.Get("/:id", handler.GetProduct)
	// bookRoutes.Put("/:id", handler.UpdateProduct)
}
