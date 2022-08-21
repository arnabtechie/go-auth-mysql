package controllers

import (
	"time"

	"github.com/arnabtechie/go-ecommerce/models"
	"github.com/arnabtechie/go-ecommerce/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {

	register := &models.User{}

	// DB := sql_connector.DB

	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
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
	register.LastLogin = time.Now()
	// err := DB.Model(&models.User{}).Create(&register)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// }
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
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
