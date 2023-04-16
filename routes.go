package main

import (
	"database/sql"
	"errors"
	"go-auth-mysql/controllers"
	"go-auth-mysql/db"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Authorization header missing",
			})
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid authorization header format",
			})
		}

		tokenString := authHeaderParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		if !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid token claims",
			})
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid user ID in token",
			})
		}
		var user db.User

		row := db.DB.QueryRow("select id, full_name, email, created_at from users where id = ?", userID)
		err = row.Scan(&user.ID, &user.FullName, &user.Email, &user.CreatedAt)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(401).JSON(fiber.Map{
					"message": "User not found",
				})
			} else {
				log.Println("Error querying database:", err)
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal server error",
				})
			}
		}

		c.Locals("user", user)

		return c.Next()
	}
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users/register", controllers.Register)
	api.Post("/users/login", controllers.Login)

	api.Use(JWTMiddleware())

	api.Get("/users/me", controllers.Profile)
}
