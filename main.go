package main

import (
	"go-auth-mysql/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success"})
	})

	db.Setup()

	app.Use(cors.New())

	SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"status": "fail", "message": "404 not found"})
	})

	app.Listen(":4000")
}
