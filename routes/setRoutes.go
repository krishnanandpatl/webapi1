package routes

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {
	app.Get("/", Hiii)
}
