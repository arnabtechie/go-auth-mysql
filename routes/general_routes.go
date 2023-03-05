package routes

import (
	"github.com/arnabtechie/go-ecommerce/controllers"
	"github.com/gofiber/fiber/v2"
)

func GenericRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.Register)

	route.Get("/logout", controllers.Logout)
}
