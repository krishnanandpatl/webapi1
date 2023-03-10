package routes

import "github.com/gofiber/fiber/v2"

func Hiii(c *fiber.Ctx) error {
	return c.SendString("hiii")
}
