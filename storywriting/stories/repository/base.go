package repository

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type StoryRepository struct {
	DB *gorm.DB
}

// NewStoryRepository initializes a new StoryRepository instance
func NewStoryRepository() *StoryRepository {
	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}

	return &StoryRepository{
		DB: db,
	}
}
