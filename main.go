package main

import (
	"golang-rest-api/config"
	"golang-rest-api/migrations"
	"golang-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Koneksi database
	config.ConnectDatabase()

	// Jalankan migrasi
	migrations.RunMigration()

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server
	app.Listen(":3000")
}
