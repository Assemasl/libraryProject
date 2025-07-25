package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	AuthorID    uint
	Author      User `gorm:"foreignKey:AuthorID"`
}
