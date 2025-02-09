package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func UpDBLogIn() {
	err := godotenv.Load()
	if err != nil {
		println("error in get value from env file")
		os.Exit(1)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")


	// בדיקה שכל המשתנים קיימים
	if host == "" || user == "" || password == "" || dbname == "" {
		fmt.Println("One or more environment variables are missing")
		os.Exit(1)
	}

	// בניית ה-DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", host, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("error in connect to DB")
		os.Exit(1)
	}

	DB = db
}
