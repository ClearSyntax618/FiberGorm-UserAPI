package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON("Buenas tardes")
}

func Login(c *fiber.Ctx) error {
	return c.Status(200).JSON("Hola Mundo")
}
