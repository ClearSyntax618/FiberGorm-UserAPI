package router

import (
	"github.com/ClearSyntax618/FiberGorm-UserAPI/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
}
