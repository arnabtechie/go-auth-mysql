package main

import (
	"go-auth-mysql/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users/register", controllers.Register)
}
