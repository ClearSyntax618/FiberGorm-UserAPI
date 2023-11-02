package main

import (
	"log"

	"github.com/ClearSyntax618/FiberGorm-UserAPI/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3031"))
}
