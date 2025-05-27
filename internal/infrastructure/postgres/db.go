package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect connects to PostgreSQL and returns the *gorm.DB instance.
// It no longer calls AutoMigrate or uses model directly — separation of concern.
func Connect() *gorm.DB {
	env := os.Getenv("APP_ENV")
	envFile := ".env"
	if env == "development" {
		envFile = ".env.development"
	}

	_ = loadEnv(envFile)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	log.Printf("✅ Connected to DB (%s)", env)
	return db
}

func loadEnv(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("❌ Failed to load env file %s: %v", file, err)
	}
	return nil
}
