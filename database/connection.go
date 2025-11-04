package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUser := getEnv("DB_USER")
	dbPass := getEnv("DB_PASS")
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")
	dbName := getEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully")
}

func getEnv(key string) string {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, relying on environment variables")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
