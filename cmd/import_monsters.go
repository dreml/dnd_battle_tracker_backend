package main

import (
	"battle_tracker/internal/monsters"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongo_client *mongo.Client
	collection   *mongo.Collection
)

type GetMonstersResponse struct {
	Count   int           `json:"count"`
	Results []MonsterInfo `json:"results"`
}

type MonsterInfo struct {
	Index string `json:"index"`
	Name  string `json:"name"`
}

type GetMonsterResponse struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Armor []struct {
		ArmorType string `json:"type"`
		Value     int    `json:"value"`
	} `json:"armor_class"`
	Image     string `json:"image"`
	HitPoints int    `json:"hit_points"`
}

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

	var response GetMonstersResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	for i, monster := range response.Results {
		importMonster(monster)
		fmt.Printf("imported %v: %d of %d\n", monster.Name, i+1, len(response.Results))
	}
}

func importMonster(monsterInfo MonsterInfo) {
	resp, err := http.Get("https://www.dnd5eapi.co/api/monsters/" + monsterInfo.Index)
	if err != nil {
		log.Fatal(err)
	}

	var r GetMonsterResponse
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatal(err)
	}

	armor := 0
	if len(r.Armor) > 0 {
		armor = r.Armor[0].Value
	}

	monster := monsters.Monster{
		ID:          primitive.NewObjectID(),
		Slug:        r.Index,
		Name:        r.Name,
		Health:      r.HitPoints,
		Armor:       armor,
		Image:       r.Image,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	_, err = collection.InsertOne(context.Background(), monster)
	if err != nil {
		log.Fatal(err)
	}
}
