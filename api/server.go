package api

import (
	"battle_tracker/pkg/common"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct{}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

func JSONHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func NewServer(listenAddr string) *http.Server {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CORSMiddleware)
	r.Use(JSONHeaderMiddleware)

	s := r.PathPrefix("/api/").Subrouter()
	s.HandleFunc("/monsters", handleGetMonsters).Methods("GET")
	s.HandleFunc("/monster/{monster}", handleGetMonster).Methods("GET")

	return &http.Server{
		Handler: r,
		Addr:    listenAddr,
	}
}

func handleGetMonsters(w http.ResponseWriter, r *http.Request) {
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

func handleGetMonster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monster := vars["monster"]

	response := map[string]string{"result": monster}
	json.NewEncoder(w).Encode(response)
	// TODO get moster
}
