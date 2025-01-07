package api

import (
	"battle_tracker/internal/campaigns"
	"battle_tracker/internal/characters"
	"battle_tracker/internal/monsters"
	"context"
	"log"

	"github.com/labstack/echo/v4"
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

	db := client.Database("battle_tracker")

	// monsters router
	monstersHandler := monsters.NewHandler(db)
	mr := e.Group("/monsters")
	mr.GET("", monstersHandler.GetMonsters)
	mr.POST("", monstersHandler.CreateMonster)
	// mr.GET("/:monsterSlug", monstersHandler.GetMonsterBySlug)
	mr.GET("/:monsterId", monstersHandler.GetMonster)
	mr.PATCH("/:monsterId", monstersHandler.UpdateMonster)
	mr.DELETE("/:monsterId", monstersHandler.DeleteMonster)

	// campaigns router
	campaignsHandler := campaigns.NewHandler(db)
	cr := e.Group("/campaigns")
	cr.GET("", campaignsHandler.GetCampaigns)
	cr.POST("", campaignsHandler.CreateCampaign)
	cr.GET("/:campaignId", campaignsHandler.GetCampaign)
	cr.PATCH("/:campaignId", campaignsHandler.UpdateCampaign)
	cr.DELETE("/:campaignId", campaignsHandler.DeleteCampaign)

	// characters router
	charactersHandler := characters.NewHandler(db)
	chr := e.Group("/characters")
	chr.GET("", charactersHandler.GetCharacters)
	chr.POST("", charactersHandler.CreateCharacter)
	chr.GET("/:characterId", charactersHandler.GetCharacter)
	chr.PATCH("/:characterId", charactersHandler.UpdateCharacter)
	chr.DELETE("/:characterId", charactersHandler.DeleteCharacter)
}
