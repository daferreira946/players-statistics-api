package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"players/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	user := os.Getenv("POSTGRES_USER")
	host := "localhost"
	port := "5432"
	dsn := os.Getenv("DATABASE_URL")

	if dsn != "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = connection.AutoMigrate(&models.Player{})

	if err != nil {
		panic(err)
	}

	DB = connection
	err = DB.AutoMigrate(&models.Player{}, &models.Score{})

	if err != nil {
		panic(err)
	}
}
