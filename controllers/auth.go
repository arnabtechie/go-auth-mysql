package controllers

import (
	"fmt"

	"github.com/arnabtechie/go-ecommerce/models"
	"github.com/arnabtechie/go-ecommerce/utils"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	register := &models.User{}

	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	fmt.Println(register.LastName)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":        register.ID,
			"uuid":      register.UUID,
			"firstName": register.FirstName,
			"lastName":  register.LastName,
		},
	})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
