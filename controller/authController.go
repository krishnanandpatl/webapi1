package controller

import (
	"weabapi/database"
	"weabapi/models"
	"weabapi/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	
	db := database.DBConn
	
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hashPassword,err:=utils.PasswordHasher(user.Password)
	if err!=nil{
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "can't generate password", "data": err})
	}

	user.Password=hashPassword

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "user created", "data": user})
}
