package controllers

import (
	"github.com/arnabtechie/go-ecommerce/models"
	"github.com/arnabtechie/go-ecommerce/sql_connector"
	"github.com/arnabtechie/go-ecommerce/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func Register(c *fiber.Ctx) error {

	register := &models.User{}

	DB := sql_connector.DB

	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": utils.ValidatorErrors(err),
		})
	}
	register.ID = uuid.NewString()

	if err := DB.Table("users").Create(&register).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err,
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":        register.ID,
			"firstName": register.FirstName,
			"lastName":  register.LastName,
			"email":     register.Email,
		},
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

	DB := sql_connector.DB

	type Result struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	result := &Result{}

	if err := DB.Table("users").Select("id", "email", "password").Where("email = ?", signIn.Email).First(result).Error; gorm.IsRecordNotFoundError(err) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err,
		})
	}

	if result.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  "Username or password incorrect",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
