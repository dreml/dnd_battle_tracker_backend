package api

import (
	"battle_tracker/pkg/common"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/monsters", s.handleGetMonsters)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetMonsters(w http.ResponseWriter, r *http.Request) {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("battle_tracker").Collection("monsters")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var monsters []common.Monster
	if err = cursor.All(context.TODO(), &monsters); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(monsters)
}
