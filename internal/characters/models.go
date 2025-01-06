package characters

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Name        string             `bson:"name, omitempty"`
	PlayerName  string             `bson:"playerName, omitempty"`
	Avatar      string             `bson:"avatar, omitempty"`
	Health      int                `bson:"health, omitempty"`
	Armor       int                `bson:"armor, omitempty"`
	CampaignId  string             `bson:"campaignId, omitempty"`
	DateCreated time.Time          `bson:"dateCreated, omitempty"`
	DateUpdated time.Time          `bson:"dateUpdated, omitempty"`
}

type CharacterResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PlayerName  string `json:"playerName"`
	Avatar      string `json:"avatar"`
	Health      int    `json:"health"`
	Armor       int    `json:"armor"`
	CampaignId  string `json:"campaignId"`
	DateCreated int    `json:"dateCreated"`
	DateUpdated int    `json:"dateUpdated"`
}

func createCharacterResponseFromCharacter(c Character) CharacterResponse {
	return CharacterResponse{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		PlayerName:  c.PlayerName,
		Avatar:      c.Avatar,
		Health:      c.Health,
		Armor:       c.Armor,
		CampaignId:  c.CampaignId,
		DateCreated: int(c.DateCreated.Unix()),
		DateUpdated: int(c.DateUpdated.Unix()),
	}
}

type CreateCharacterRequest struct {
	Name       string `json:"name"`
	PlayerName string `json:"playerName"`
	Avatar     string `json:"avatar"`
	Health     int    `json:"health"`
	Armor      int    `json:"armor"`
	CampaignId string `json:"campaignId"`
}

type UpdateCharacterRequest struct {
	Name       *string `json:"name"`
	PlayerName *string `json:"playerName"`
	Avatar     *string `json:"avatar"`
	Health     *int    `json:"health"`
	Armor      *int    `json:"armor"`
	CampaignId *string `json:"campaignId"`
}

type GetCharactersResponse []CharacterResponse

type GetCharacterResponse CharacterResponse

type CreateCharacterResponse CharacterResponse

type UpdateCharacterResponse CharacterResponse
