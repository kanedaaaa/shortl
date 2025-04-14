package service

import (
	"fmt"

	"github.com/kanedaaaa/shortl/internal/db"
	"github.com/kanedaaaa/shortl/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(email, username, password string) error {
	var existingUser models.User

	if err := db.DB.Where("email = ?").First(&existingUser).Error; err == nil {
		return fmt.Errorf("email is already taken")
	}

	if err := db.DB.Where("username = ?").First(&existingUser).Error; err == nil {
		return fmt.Errorf("username is already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to has password: %v", err)
	}

	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
