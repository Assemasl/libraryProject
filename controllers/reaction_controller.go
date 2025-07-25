package controllers

import (
	"book-library/models"
	"book-library/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReactionInput struct {
	Type string `json:"type"` // "like" или "dislike"
}

func ReactToBook(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	bookID, _ := strconv.Atoi(c.Params("id"))

	var input ReactionInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.Type != "like" && input.Type != "dislike" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Type must be like or dislike"})
	}

	var existing models.Reaction
	if err := models.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		First(&existing).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You have already reacted"})
	}

	reaction := models.Reaction{
		UserID: userID,
		BookID: uint(bookID),
		Type:   input.Type,
	}

	if err := models.DB.Create(&reaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add reaction"})
	}

	var book models.Book
	if err := models.DB.Preload("Author").First(&book, bookID).Error; err == nil && input.Type == "like" {
		go utils.SendEmail(
			"zhaksilikova.a05@gmail.com",
			"Ваша книга получила лайк",
			"Поздравляем! Кто-то поставил лайк вашей книге: "+book.Title,
		)
	}

	return c.JSON(fiber.Map{"message": "Reaction added"})
}

func RemoveReaction(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	reactionID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reaction ID"})
	}

	var reaction models.Reaction
	if err := models.DB.First(&reaction, reactionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reaction not found"})
	}

	if reaction.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not allowed to delete this reaction"})
	}

	if err := models.DB.Delete(&reaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete reaction"})
	}

	return c.JSON(fiber.Map{"message": "Reaction deleted"})
}
