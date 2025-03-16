package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	db *gorm.DB
	once sync.Once
)

// ConnectDB returns a database connection
func ConnectDB() *gorm.DB {
	once.Do(func ()  {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			os.Getenv("DB_ADDRESS"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
		// Check if there is an error
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		log.Println("Database connected")
	})
	return db
}