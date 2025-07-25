package middlewares

import (
	"book-library/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("user_id", claims["user_id"])
		c.Locals("is_author", claims["is_author"])
		return c.Next()
	}
}

func AuthorOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAuthor, ok := c.Locals("is_author").(bool)
		if !ok || !isAuthor {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only authors allowed"})
		}
		return c.Next()
	}
}

func ReaderOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAuthor, ok := c.Locals("is_author").(bool)
		if !ok || isAuthor {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only readers allowed"})
		}
		return c.Next()
	}
}
