package utils

import "github.com/gofiber/fiber/v2"

func JSONOK(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}
