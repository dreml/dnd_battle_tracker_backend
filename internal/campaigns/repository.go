package campaigns

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignsRepository struct {
	collection *mongo.Collection
}

func NewCampaignsRepository(db *mongo.Database) *CampaignsRepository {
	return &CampaignsRepository{
		collection: db.Collection("campaigns"),
	}
}

func (cr *CampaignsRepository) findCampaigns() ([]Campaign, error) {
	cursor, err := cr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var campaigns []Campaign
	if err = cursor.All(context.TODO(), &campaigns); err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (cr *CampaignsRepository) getCampaignByID(campaignId string) (Campaign, error) {
	id, _ := primitive.ObjectIDFromHex(campaignId)
	filter := bson.M{"_id": id}

	var campaign Campaign
	err := cr.collection.FindOne(context.Background(), filter).Decode(&campaign)
	if err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignsRepository) createCampaign(r CreateCampaignRequest) (Campaign, error) {
	campaign := Campaign{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	_, err := cr.collection.InsertOne(context.Background(), campaign)
	if err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignsRepository) updateCampaign(campaignId string, r UpdateCampaignRequest) (Campaign, error) {
	id, _ := primitive.ObjectIDFromHex(campaignId)

	var campaign Campaign
	filter := bson.M{"_id": id}
	err := cr.collection.FindOne(context.Background(), filter).Decode(&campaign)
	if err != nil {
		return Campaign{}, err
	}

	campaign.Name = r.Name
	campaign.DateUpdated = time.Now()

	_, err = cr.collection.UpdateByID(context.Background(), id, bson.M{"$set": campaign})
	if err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}

func (cr *CampaignsRepository) deleteCampaign(campaignId string) error {
	id, _ := primitive.ObjectIDFromHex(campaignId)
	filter := bson.M{"_id": id}

	_, err := cr.collection.DeleteOne(context.Background(), filter)

	return err
}
