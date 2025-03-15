package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse untuk response sukses
func SuccessResponse(message string, data interface{}) fiber.Map {
	return fiber.Map{
		"message": message,
		"error":   false,
		"data":    data,
	}
}

// ErrorResponse untuk response error
func ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		"message": message,
		"error":   true,
		"data":    nil,
	}
}
