package main

import (
	"log"

	"github.com/arnabtechie/go-ecommerce/middlewares"
	"github.com/arnabtechie/go-ecommerce/routes"
	connection "github.com/arnabtechie/go-ecommerce/sql_connector"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := connection.InitDB()
	if err != nil {
		panic(err.Error())
	}

	defer connection.DB.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})
	middlewares.FiberMiddleware(app)
	routes.GenericRoutes(app)
	routes.NotFoundRoute(app)
	log.Fatal(app.Listen(":4000"))
}
