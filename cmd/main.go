package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/theborzet/library_project/internal/db/init"
	"github.com/theborzet/library_project/internal/handlers"
	"github.com/theborzet/library_project/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Some problems with config", err)
	}
	app := fiber.New()

	db := db.Init(config)
	handlers.RegistrationRoutess(app, db)

	app.Listen(config.Port)
}
