package middlewares

import (
	"golang-rest-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(utils.ErrorResponse("Unauthorized"))
	}

	// Ambil token dari header
	tokenString := strings.Split(authHeader, " ")[1]
	claims := jwt.MapClaims{}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecretkey"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(utils.ErrorResponse("Invalid token"))
	}

	// Simpan userID ke `Locals`
	userID := uint(claims["id"].(float64))
	c.Locals("userID", userID)

	return c.Next()
}
