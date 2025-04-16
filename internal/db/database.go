package db

import (
	"log"
	"os"

	"github.com/kanedaaaa/shortl/internal/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}

	DB = db

	logrus.Info("Database connection established")

	err = db.AutoMigrate(&models.User{}, &models.Link{})
	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}

	logrus.Info("Database migrated successfully")

}
