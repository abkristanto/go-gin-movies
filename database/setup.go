package database

import (
	"fmt"
	"log"
	"os"

	"github.com/abkristanto/go-gin-movies/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		HOST        = os.Getenv("HOST")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.Genre{})
	db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	return nil
}

func Get() *gorm.DB {
	return db
}
