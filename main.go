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
		log.Println("Error loading .env file", err.Error())
	}
	config.ConnectDatabase()
	routes.HandleRequests()
}
