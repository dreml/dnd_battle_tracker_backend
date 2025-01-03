package campaigns

import (
	"battle_tracker/pkg/common"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignService struct {
	campaignsCollection *mongo.Collection
}

func NewCampaignService(d *mongo.Database) *CampaignService {
	return &CampaignService{
		campaignsCollection: d.Collection("campaigns"),
	}
}

func (cs *CampaignService) HandleGetCampaigns(c echo.Context) error {
	cursor, err := cs.campaignsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var campaigns []Campaign
	if err = cursor.All(context.TODO(), &campaigns); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (cs *CampaignService) HandleGetCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")
	filter := bson.M{"id": campaignId}

	var campaign Campaign
	err := cs.campaignsCollection.FindOne(context.Background(), filter).Decode(&campaign)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	} else {
		return c.JSON(http.StatusOK, campaign)
	}
}

func (cs *CampaignService) HandleCreateCampaign(c echo.Context) error {
	var campaign Campaign

	c.Bind(&campaign)
	r, err := cs.campaignsCollection.InsertOne(context.Background(), campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	campaign.ID = r.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusOK, campaign)
}

func (cs *CampaignService) HandleUpdateCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")

	var updateData Campaign
	c.Bind(&updateData)
	if updateData.Name == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	_, err := cs.campaignsCollection.UpdateByID(context.Background(), campaignId, updateData)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	var campaign Campaign
	filter := bson.M{"id": campaignId}
	err = cs.campaignsCollection.FindOne(context.Background(), filter).Decode(&campaign)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	} else {
		return c.JSON(http.StatusOK, campaign)
	}
}

func (cs *CampaignService) HandleDeleteCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")
	filter := bson.M{"id": campaignId}

	var campaign Campaign
	err := cs.campaignsCollection.FindOne(context.Background(), filter).Decode(&campaign)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	_, err = cs.campaignsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, nil)
}
