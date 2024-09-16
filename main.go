package main

import (
	"battle_tracker/api"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	apiRouter := api.NewApiRouter()
	// socketServer := socketio.NewServer(nil)

	http.Handle("/api/", apiRouter)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}
