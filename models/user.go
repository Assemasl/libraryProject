package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string
	IsAuthor     bool
	Books        []Book `gorm:"foreignKey:AuthorID"` // внешний key
}
