package campaigns

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Name        string             `bson:"name, omitempty"`
	DateCreated time.Time          `bson:"dateCreated, omitempty"`
	DateUpdated time.Time          `bson:"dateUpdated, omitempty"`
}

type CampaignResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateCreated int    `json:"dateCreated"`
	DateUpdated int    `json:"dateUpdated"`
}

func createCampaignResponseFromCampaign(c Campaign) CampaignResponse {
	return CampaignResponse{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		DateCreated: int(c.DateCreated.Unix()),
		DateUpdated: int(c.DateUpdated.Unix()),
	}
}

type CreateCampaignRequest struct {
	Name string `json:"name"`
}

type UpdateCampaignRequest struct {
	Name string `json:"name"`
}

type GetCampaignsResponse []CampaignResponse

type GetCampaignResponse CampaignResponse

type CreateCampaignResponse CampaignResponse

type UpdateCampaignResponse CampaignResponse
