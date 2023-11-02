package router

import (
	"github.com/ClearSyntax618/FiberGorm-UserAPI/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.HelloWorld)
	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
}
