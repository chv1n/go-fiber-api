package product

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewProductController(s Service) *Controller {
	return &Controller{s}
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	products, err := c.service.GetProducts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(products)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	var body Product

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := c.service.CreateProduct(&body); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(body)
}
