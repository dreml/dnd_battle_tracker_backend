package monsters

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monster struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Slug        string             `bson:"slug, omitempty"`
	Name        string             `bson:"name, omitempty"`
	Health      int                `bson:"health, omitempty"`
	Armor       int                `bson:"armor, omitempty"`
	Image       string             `bson:"image, omitempty"`
	DateCreated time.Time          `bson:"dateCreated, omitempty"`
	DateUpdated time.Time          `bson:"dateUpdated, omitempty"`
}

type MonsterInfo struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type MonsterResponse struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Health      int    `json:"health"`
	Armor       int    `json:"armor"`
	Image       string `json:"image"`
	DateCreated int    `json:"dateCreated"`
	DateUpdated int    `json:"dateUpdated"`
}

func createMonsterInfoFromMonster(m Monster) MonsterInfo {
	return MonsterInfo{
		ID:   m.ID.Hex(),
		Name: m.Name,
		Slug: m.Slug,
	}
}

func createMonsterResponseFromMonster(m Monster) MonsterResponse {
	return MonsterResponse{
		ID:          m.ID.Hex(),
		Slug:        m.Slug,
		Name:        m.Name,
		Image:       m.Image,
		Health:      m.Health,
		Armor:       m.Armor,
		DateCreated: int(m.DateCreated.Unix()),
		DateUpdated: int(m.DateUpdated.Unix()),
	}
}

type CreateMonsterRequest struct {
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Health int    `json:"health"`
	Armor  int    `json:"armor"`
}

type UpdateMonsterRequest struct {
	Slug   *string `json:"slug"`
	Name   *string `json:"name"`
	Image  *string `json:"image"`
	Health *int    `json:"health"`
	Armor  *int    `json:"armor"`
}

type GetMonstersResponse []MonsterInfo

type GetMonsterResponse MonsterResponse

type GetMonsterByResponse MonsterResponse

type CreateMonsterResponse MonsterResponse

type UpdateMonsterResponse MonsterResponse
