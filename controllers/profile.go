package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Ambil ID user dari token JWT

	// Cek apakah profile sudah ada
	var existingProfile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&existingProfile).Error; err == nil {
		return c.Status(400).JSON(utils.ErrorResponse("Profile already exists"))
	}

	// Parse input request
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Invalid request"))
	}

	// Set user ID dari token
	profile.UserID = userID

	// Simpan profile ke database
	if err := config.DB.Create(&profile).Error; err != nil {
		return c.Status(500).JSON(utils.ErrorResponse("Failed to create profile"))
	}

	return c.JSON(utils.SuccessResponse("Profile created successfully", profile))
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint) // Ambil user ID dari token JWT

	var profile models.Profile
	result := config.DB.Where("user_id = ?", userID).First(&profile)

	// Jika profile tidak ditemukan
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(utils.ErrorResponse("Profile not found"))
	}

	return c.JSON(utils.SuccessResponse("Profile retrieved successfully", profile))
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint) // Ambil user ID dari token JWT

	var profile models.Profile
	result := config.DB.Where("user_id = ?", userID).First(&profile)

	// Jika profile tidak ditemukan
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(utils.ErrorResponse("Profile not found"))
	}

	// Update profile berdasarkan data yang dikirim
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Invalid request"))
	}

	config.DB.Save(&profile)

	return c.JSON(utils.SuccessResponse("Profile updated successfully", profile))
}

func DeleteProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint) // Ambil user ID dari token JWT

	var profile models.Profile
	result := config.DB.Where("user_id = ?", userID).First(&profile)

	// Jika profile tidak ditemukan
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(utils.ErrorResponse("Profile not found"))
	}

	config.DB.Delete(&profile)

	return c.JSON(utils.SuccessResponse("Profile deleted successfully", nil))
}
