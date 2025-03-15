package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoutes(app *fiber.App) {
	profile := app.Group("/profile", middlewares.AuthMiddleware)

	profile.Post("/", controllers.CreateProfile)
	profile.Get("/", controllers.GetProfile)
	profile.Put("/", controllers.UpdateProfile)
	profile.Delete("/", controllers.DeleteProfile)
}
