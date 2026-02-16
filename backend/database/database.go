package database

import (
	"fmt"
	"log"
	"os"
	"trip_splitter/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 🚀 disables prepared statement caching
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	// Auto migrate the schema
	DB.AutoMigrate(&models.User{},&models.Trip{},&models.Expense{},&models.ExpenseSplit{},&models.Debt{})
	fmt.Println(" DB connected successfully....")
}
