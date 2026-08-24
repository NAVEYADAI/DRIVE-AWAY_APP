package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// .env is optional: in Docker/CI the DB_* vars are usually injected directly
	// into the process environment, so a missing file here is not fatal.
	if err := godotenv.Load(); err != nil {
		println("no .env file found, reading DB config from the environment")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "require"
	}

	// בדיקה שכל המשתנים קיימים
	if host == "" || user == "" || password == "" || dbname == "" {
		fmt.Println("One or more environment variables are missing")
		os.Exit(1)
	}

	// בניית ה-DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("error in connect to DB")
		os.Exit(1)
	}

	DB = db
}
