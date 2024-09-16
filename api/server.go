package api

import (
	"battle_tracker/pkg/common"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApiServer struct {
	monstersCollection *mongo.Collection
}

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

func NewApiRouter() *mux.Router {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	monstersCollection := client.Database("battle_tracker").Collection("monsters")
	as := &ApiServer{
		monstersCollection: monstersCollection,
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CORSMiddleware)
	r.Use(JSONHeaderMiddleware)

	s := r.PathPrefix("/api/").Subrouter()
	s.HandleFunc("/monsters", as.handleGetMonsters).Methods("GET")
	s.HandleFunc("/monster/{monster}", as.handleGetMonster).Methods("GET")

	return r
}

func (as *ApiServer) handleGetMonsters(w http.ResponseWriter, r *http.Request) {
	cursor, err := as.monstersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var monsters []common.MonsterInfo
	if err = cursor.All(context.TODO(), &monsters); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(monsters)
}

func (as *ApiServer) handleGetMonster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monster_index := vars["monster"]

	filter := bson.M{"index": monster_index}

	var monster common.Monster
	err := as.monstersCollection.FindOne(context.Background(), filter).Decode(&monster)

	if errors.Is(err, mongo.ErrNoDocuments) {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(monster)
	}
}
