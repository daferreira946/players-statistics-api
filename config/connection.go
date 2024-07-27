package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"players/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	usingUrl := true
	log.Println(dsn)

	if dsn == "" {
		database := os.Getenv("POSTGRES_DB")
		password := os.Getenv("POSTGRES_PASSWORD")
		user := os.Getenv("POSTGRES_USER")
		host := "localhost"
		port := "5432"
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
		usingUrl = false
	}

	connection, err := connect(dsn, usingUrl)

	if err != nil {
		panic(err)
	}

	DB = connection
	err = DB.AutoMigrate(&models.Player{}, &models.Score{})

	if err != nil {
		panic(err)
	}
}

func connect(dsn string, usingUrl bool) (*gorm.DB, error) {
	if !usingUrl {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	sqlDB, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return gorm.Open(postgres.New(
		postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
}
