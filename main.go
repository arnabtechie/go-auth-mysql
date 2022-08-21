package main

import (
	"log"

	"github.com/arnabtechie/go-ecommerce/middlewares"
	"github.com/arnabtechie/go-ecommerce/routes"
	"github.com/arnabtechie/go-ecommerce/sql_connector"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	middlewares.FiberMiddleware(app)
	sql_connector.Connection()
	routes.GenericRoutes(app)
	routes.NotFoundRoute(app)
	log.Fatal(app.Listen(":4000"))
}
