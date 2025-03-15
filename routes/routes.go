package routes

import "github.com/gofiber/fiber/v2"

// SetupRoutes memanggil semua routes
func SetupRoutes(app *fiber.App) {
	UserRoutes(app)
	ProfileRoutes(app)
}
