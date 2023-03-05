package controllers

import (
	"log"

	connection "github.com/arnabtechie/go-ecommerce/sql_connector"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {

	type User struct {
		UUID     string `json:"uuid"`
		FullName string `json:"full_name" validate:"required,lte=255"`
		Email    string `json:"email" validate:"required,email,lte=255"`
		Password string `json:"password" validate:"required,lte=255"`
	}

	register := &User{}

	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	v := validator.New()

	if err := v.Struct(register); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	rows, err := connection.DB.Query("select * from users where email=?", register.Email)
	if err != nil {
		return err
	}
	defer rows.Close()

	register.UUID = uuid.NewString()
	log.Println(*register)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
	})
}

func Login(c *fiber.Ctx) error {

	type SignIn struct {
		Email    string `json:"email" validate:"required,email,lte=255"`
		Password string `json:"password" validate:"required,lte=255"`
	}

	signIn := &SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	v := validator.New()

	if err := v.Struct(signIn); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    "User logged out",
	})
}
