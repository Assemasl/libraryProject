package models

type Reaction struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	BookID uint
	Type   string // "like" или "dislike"
}
