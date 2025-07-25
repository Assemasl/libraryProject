package controllers

import (
	"book-library/models"
	"book-library/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	books, err := repository.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch books"})
	}
	return c.JSON(books)
}

func GetBookByID(c *fiber.Ctx) error {
	bookID, _ := strconv.Atoi(c.Params("id"))

	book, err := repository.GetBookByID(uint(bookID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	likes, _ := repository.CountReactions(uint(bookID), "like")
	dislikes, _ := repository.CountReactions(uint(bookID), "dislike")

	return c.JSON(fiber.Map{
		"book":     book,
		"likes":    likes,
		"dislikes": dislikes,
	})
}

func CreateBook(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		AuthorID:    uint(userID),
	}

	if err := repository.CreateBook(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	bookID, _ := strconv.Atoi(c.Params("id"))
	userID := uint(c.Locals("user_id").(float64))

	book, err := repository.GetBookByID(uint(bookID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if book.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can only update your own books"})
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	book.Title = input.Title
	book.Description = input.Description

	if err := repository.UpdateBook(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	bookID, _ := strconv.Atoi(c.Params("id"))
	userID := uint(c.Locals("user_id").(float64))

	book, err := repository.GetBookByID(uint(bookID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if book.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can only delete your own books"})
	}

	if err := repository.DeleteBook(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.JSON(fiber.Map{"message": "Book deleted"})
}
