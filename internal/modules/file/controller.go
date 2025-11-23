package file

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewFileController(s Service) *Controller {
	return &Controller{s}
}

func (c *Controller) Upload(ctx *fiber.Ctx) error {
	// Get file from request
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to get file from request",
		})
	}

	// Upload file
	filePath, err := c.service.UploadFile(ctx, file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"path":    filePath,
	})
}
