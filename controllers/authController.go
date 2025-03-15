package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Invalid request"))
	}

	// Cek apakah email sudah digunakan
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(utils.ErrorResponse("Email already exists"))
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(utils.ErrorResponse("Failed to register user"))
	}

	// Buat profile default untuk user
	profile := &models.Profile{
		UserID:   user.ID,
		Username: user.Name,
		Bio:      "No bio yet",
		Phone:    "",
		Image:    "default.png",
	}
	config.DB.Create(&profile)

	// Assign profile ke user agar dikembalikan dalam response
	user.Profile = profile

	return c.JSON(utils.SuccessResponse("User registered successfully", user))
}

func Login(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(utils.ErrorResponse("Invalid request"))
	}

	// Cari user berdasarkan email, termasuk profile-nya
	config.DB.Preload("Profile").Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(401).JSON(utils.ErrorResponse("Invalid credentials"))
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("mysecretkey"))

	return c.JSON(utils.SuccessResponse("Login successful", fiber.Map{
		"token": signedToken,
		"user":  user,
	}))
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	// Query semua user dari database
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(utils.ErrorResponse("Failed to retrieve users"))
	}

	// Jika tidak ada user
	if len(users) == 0 {
		return c.JSON(utils.SuccessResponse("No users found", []models.User{}))
	}

	return c.JSON(utils.SuccessResponse("Users retrieved successfully", users))
}

func GetUserById(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL
	id := c.Params("id")

	var user models.User

	// Gunakan Preload untuk memuat Profile
	result := config.DB.Preload("Profile").First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(utils.ErrorResponse("User not found"))
	}

	return c.JSON(utils.SuccessResponse("User retrieved successfully", user))
}
