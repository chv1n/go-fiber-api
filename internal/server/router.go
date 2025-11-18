package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-fiber-basic-api/internal/modules/user"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	RegisterMiddleware(app)

	api := app.Group("/api")

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserController(userService)

	userRoute := api.Group("/users")
	{
		userRoute.Get("/", userHandler.GetAll)
		userRoute.Post("/", userHandler.Create)
	}
}
