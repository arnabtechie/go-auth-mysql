package controllers

import (
	"errors"

	"github.com/arnabtechie/go-ecommerce/models"
	"github.com/arnabtechie/go-ecommerce/sql_connector"
	"github.com/arnabtechie/go-ecommerce/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	signIn := &models.SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	DB := sql_connector.DB
	result := &models.User{}

	err := DB.Table("users").Where("email = ?", signIn.Email).First(result)

	if err.Error != nil || errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
