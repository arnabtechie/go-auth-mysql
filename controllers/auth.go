package controllers

import (
	"github.com/arnabtechie/go-ecommerce/connection"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {

	type User struct {
		ID       string `json:"id"`
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

	// rows, err := connection.DB.Query("select * from users where email = ?", register.Email)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"success": false,
	// 		"errors":  err.Error(),
	// 	})
	// }
	// defer rows.Close()

	// log.Println(rows)

	// var users []User

	// for rows.Next() {
	// 	var user User
	// 	err = rows.Scan(&user.ID, &user.UUID, &user.FullName, &user.Email, &user.Password)
	// 	if err != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"success": false,
	// 			"errors":  err.Error(),
	// 		})
	// 	}
	// 	users = append(users, user)
	// }

	// log.Println(users)

	register.UUID = uuid.NewString()

	result, err := connection.DB.Exec("insert into users (full_name, email, password, uuid) values (?, ?, ?, ?)", register.FullName, register.Email, register.Password, register.UUID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    rowsAffected,
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
			"success": false,
			"errors":  err.Error(),
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
		"message": "User logged out",
	})
}
