package main

import (
	"github.com/joho/godotenv"
	"log"
	"players/config"
	"players/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	routes.HandleRequests()
}
