package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterMiddleware(app *fiber.App) {
	app.Use(logger.New())
}
