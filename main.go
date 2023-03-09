package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// fmt.Println("hiii")
		return c.SendString("hiii")
	})
	app.Listen(":3000")
}
