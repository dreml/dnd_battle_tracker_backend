package main

import (
	"battle_tracker/pkg/common"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongo_client *mongo.Client
	collection   *mongo.Collection
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	mongo_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	collection = mongo_client.Database("battle_tracker").Collection("monsters")

	resp, err := http.Get("https://www.dnd5eapi.co/api/monsters")
	if err != nil {
		log.Fatal(err)
	}

	var response common.MonsterInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	for i, monster := range response.Results {
		importMonster(monster)
		fmt.Printf("imported %v: %d of %d\n", monster.Name, i+1, len(response.Results))
	}
}

func importMonster(monsterInfo common.MonsterInfo) {
	resp, err := http.Get("https://www.dnd5eapi.co/api/monsters/" + monsterInfo.Index)
	if err != nil {
		log.Fatal(err)
	}

	var monster common.Monster
	if err = json.NewDecoder(resp.Body).Decode(&monster); err != nil {
		log.Fatal(err)
	}

	_, err = collection.InsertOne(context.Background(), monster)
	if err != nil {
		log.Fatal(err)
	}
}
