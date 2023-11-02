package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	fmt.Println("Hola Mundo")
	return c.Status(200).JSON("Hola Mundo")
}
