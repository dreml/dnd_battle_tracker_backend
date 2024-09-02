package main

import (
	"battle_tracker/api"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server := api.NewServer(":3000")
	log.Fatal(server.Start())
}