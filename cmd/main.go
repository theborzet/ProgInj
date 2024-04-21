package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
	db "github.com/theborzet/library_project/internal/db/init"
	"github.com/theborzet/library_project/internal/handlers"
	"github.com/theborzet/library_project/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Some problems with config", err)
	}
	engine := html.New("./internal/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db := db.Init(config)
	handlers.RegistrationRoutess(app, db, engine)

	app.Listen(config.Port)
}
