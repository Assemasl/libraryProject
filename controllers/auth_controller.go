package controllers

import (
	"book-library/models"
	"book-library/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAuthor bool   `json:"is_author"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
		IsAuthor:     input.IsAuthor,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(fiber.Map{"message": "Registration successful :_("})
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	token, err := utils.GenerateJWT(user.ID, user.IsAuthor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
