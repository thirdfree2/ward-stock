package database

import (
	"fmt"
	"log"
	"os"
	"ward-stock-backend/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	env := os.Getenv("APP_ENV")
	envFile := ".env.development"

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Failed to load env file %s: %v", envFile, err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Printf("Connected to DB (%s)\n", env)

	DB.AutoMigrate(&model.User{})
}
