package user

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewUserController(s Service) *Controller {
	return &Controller{s}
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	users, err := c.service.GetUsers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(users)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	var body User

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := c.service.CreateUser(&body)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(body)
}
