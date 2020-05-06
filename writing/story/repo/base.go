package repo

import (
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type StoryRepo struct {
	db *gorm.DB
}

// NewStoryRepo initializes a new StoryRepo instance
func NewStoryRepo() *StoryRepo {
	return &StoryRepo{
		db: getDB(),
	}
}

var once sync.Once
var singletonDB *gorm.DB

// getDB ensures database connection is only opened once in the beginning
func getDB() *gorm.DB {
	once.Do(func() {
		var err error
		singletonDB, err = gorm.Open("mysql", getConnectionString())
		if err != nil {
			log.Fatalf("Database connection failed: %s", err.Error())
		}
	})
	return singletonDB
}

func getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
}
