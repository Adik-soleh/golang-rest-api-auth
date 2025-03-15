package routes

import (
	"golang-rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	app.Get("/users", controllers.GetAllUsers)
	app.Get("/users/:id", controllers.GetUserById)
}
