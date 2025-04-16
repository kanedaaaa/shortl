package db

import (
	"github.com/kanedaaaa/shortl/internal/app/errors"
	"github.com/kanedaaaa/shortl/internal/db/models"
)

func GetUserOrPanic(userID uint) (models.User, *errors.CustomError) {
	var user models.User

	if err := DB.Where("ID = ?", userID).First(&user).Error; err != nil {
		return user, errors.NotFoundError("couldnt find user with given id")
	}

	return user, nil
}
