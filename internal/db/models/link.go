package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model

	OriginalURL  string `gorm:"not null"`
	ShortenedURL string `gorm:"not null"`
	UserID       uint   `gorm:"not null"`
}
