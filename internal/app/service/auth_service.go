package service

import (
	"github.com/kanedaaaa/shortl/internal/app/errors"
	"github.com/kanedaaaa/shortl/internal/db"
	"github.com/kanedaaaa/shortl/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(email, username, password string) *errors.CustomError {
	var existingUser models.User

	if err := db.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.ConflictError("email is already taken")
	}

	if err := db.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errors.ConflictError("username is already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServerError(err)
	}

	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return errors.InternalServerError(err)
	}

	return nil
}
