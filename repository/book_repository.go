package repository

import (
	"book-library/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := models.DB.Preload("Author").Find(&books).Error
	return books, err
}

func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	err := models.DB.Preload("Author").Preload("Reactions").First(&book, id).Error
	return book, err
}

func CreateBook(book *models.Book) error {
	return models.DB.Create(book).Error
}

func UpdateBook(book *models.Book) error {
	return models.DB.Save(book).Error
}

func DeleteBook(book *models.Book) error {
	return models.DB.Delete(book).Error
}

func CountReactions(bookID uint, reactionType string) (int64, error) {
	var count int64
	err := models.DB.Model(&models.Reaction{}).
		Where("book_id = ? AND type = ?", bookID, reactionType).
		Count(&count).Error
	return count, err
}
