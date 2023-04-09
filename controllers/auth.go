package controllers

import (
	"log"
	"os"

	"go-auth-mysql/db"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		FullName string `json:"full_name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	req := new(RegisterRequest)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid request body",
			"errors":  err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	row := db.DB.QueryRow("select count(*) from users where email = ?", req.Email)

	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println(err)
	}

	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Email already exists, try logging in"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"errors": err.Error()})
	}

	result, err := db.DB.Exec("insert into users (full_name, email, password) values (?, ?, ?)", req.FullName, req.Email, hashedPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to insert user into database",
			"errors":  err.Error(),
		})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get user ID",
			"errors":  err.Error(),
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	secretKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create JWT token",
			"errors":  err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":        id,
		"email":     req.Email,
		"full_name": req.FullName,
		"token":     tokenString,
	})
}
