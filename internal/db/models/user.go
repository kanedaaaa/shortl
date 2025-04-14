package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Links    []Link `gorm:"foreignKey:UserID"`
}
