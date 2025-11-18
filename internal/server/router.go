package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-fiber-basic-api/internal/modules/user"
	"go-fiber-basic-api/internal/modules/product"   // ← เพิ่มตรงนี้
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	RegisterMiddleware(app)

	api := app.Group("/api")

	// USER
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserController(userService)

	userRoute := api.Group("/users")
	userRoute.Get("/", userHandler.GetAll)
	userRoute.Post("/", userHandler.Create)

	// PRODUCT
	productRepo := product.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductController(productService)

	productRoute := api.Group("/products")
	productRoute.Get("/", productHandler.GetAll)
	productRoute.Post("/", productHandler.Create)
}
