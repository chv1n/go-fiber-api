package file

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	UploadFile(ctx *fiber.Ctx, file *multipart.FileHeader) (string, error)
}

type service struct{}

func NewFileService() Service {
	return &service{}
}

func (s *service) UploadFile(ctx *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	// Ensure uploads directory exists
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create upload directory: %w", err)
		}
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	// Save file
	if err := ctx.SaveFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filePath, nil
}
