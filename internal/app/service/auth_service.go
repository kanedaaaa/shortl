package service

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

func Login(email, password string) (string, *errors.CustomError) {
	var user models.User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return "", errors.NotFoundError("user with given email can not be found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.AuthError("wrong password provided")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.InternalServerError(err)
	}

	return signedToken, nil
}
