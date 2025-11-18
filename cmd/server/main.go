package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-fiber-basic-api/internal/config"
	"go-fiber-basic-api/internal/database"
	"go-fiber-basic-api/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg)

	app := fiber.New()

	server.SetupRoutes(app, db)

	log.Println("Server running on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
