package api

import (
	"battle_tracker/pkg/common"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApiServer struct {
	monstersCollection *mongo.Collection
}

func NewEchoRouter(e *echo.Group) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	monstersCollection := client.Database("battle_tracker").Collection("monsters")
	as := &ApiServer{
		monstersCollection: monstersCollection,
	}
	e.GET("/monsters", as.handleGetMonsters)
	e.GET("/monsters/:monster_index", as.handleGetMonster)
}

func (as *ApiServer) handleGetMonsters(c echo.Context) error {
	cursor, err := as.monstersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var monsters []common.MonsterInfo
	if err = cursor.All(context.TODO(), &monsters); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, monsters)
}

func (as *ApiServer) handleGetMonster(c echo.Context) error {
	monster_index := c.Param("monster_index")
	filter := bson.M{"index": monster_index}

	var monster common.Monster
	err := as.monstersCollection.FindOne(context.Background(), filter).Decode(&monster)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, monster)
	}
}
