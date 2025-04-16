package service

import (
	"github.com/kanedaaaa/shortl/internal/app/errors"
	"github.com/kanedaaaa/shortl/internal/db"
	"github.com/kanedaaaa/shortl/internal/db/models"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortenURL(userID uint, link string) (string, *errors.CustomError) {
	var user models.User

	if err := db.DB.Where("ID = ?", userID).First(&user).Error; err != nil {
		return "", errors.NotFoundError("couldnt find user with given id")
	}

	newLink := models.Link{
		OriginalURL: link,
		UserID:      userID,
	}

	if err := db.DB.Create(&newLink).Error; err != nil {
		return "", errors.InternalServerError(err)
	}

	shortened := encodeBase62(newLink.ID)

	if err := db.DB.Model(&newLink).Update("ShortenedURL", shortened).Error; err != nil {
		return "", errors.InternalServerError(err)
	}

	return shortened, nil
}

func GetURL(userID uint) ([]models.Link, *errors.CustomError) {
	var urls []models.Link

	if err := db.DB.Where("user_id = ?", userID).Find(&urls).Error; err != nil {
		return urls, errors.NotFoundError("couldnt find any URLs")
	}

	return urls, nil
}

func encodeBase62(num uint) string {
	if num == 0 {
		return "0"
	}

	var result []byte

	for num > 0 {
		remainder := num % 62
		result = append([]byte{base62Chars[remainder]}, result...)
		num = num / 62
	}

	return string(result)
}
