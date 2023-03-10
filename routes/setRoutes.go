package routes

import (
	"weabapi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", Hiii)

	api := app.Group("/api", logger.New())

	//auth
	auth := api.Group("/auth")
	auth.Post("/register", controller.RegisterUser)
}
