package service

import (
	"time"

	"github.com/kanedaaaa/shortl/internal/app/errors"
	"github.com/kanedaaaa/shortl/internal/db"
	"github.com/kanedaaaa/shortl/internal/db/models"
)

type SafeUser struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUser(userID uint) (SafeUser, *errors.CustomError) {
	user, err := db.GetUserOrPanic(userID)
	if err != nil {
		return SafeUser{}, err
	}

	safeUser := toSafeUser(user)
	return safeUser, nil
}

func toSafeUser(user models.User) SafeUser {
	return SafeUser{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
