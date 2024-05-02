package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html/v2"
	db "github.com/theborzet/library_project/internal/db/init"
	"github.com/theborzet/library_project/internal/handlers"
	"github.com/theborzet/library_project/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Some problems with config", err)
	}
	engine := html.New("./internal/templates", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(csrf.New(csrf.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			accepts := c.Accepts("html", "json")
			path := c.Path()
			if accepts == "json" || strings.HasPrefix(path, "/api/") {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Forbidden",
				})
			}
			return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
				"Title":  "Forbidden",
				"Status": fiber.StatusForbidden,
			}, "layouts/main")
		},
	}))

	app.Static("/static", "./internal/static")

	db := db.Init(config)
	handlers.RegistrationRoutess(app, db)

	app.Listen(config.Port)
}
