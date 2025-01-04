package campaigns

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignService struct {
	campaignsCollection *mongo.Collection
	repository          *CampaignsRepository
}

func NewCampaignService(db *mongo.Database) *CampaignService {
	return &CampaignService{
		repository: NewCampaignsRepository(db),
	}
}

func (cs *CampaignService) getAllCampaigns() ([]Campaign, error) {
	return cs.repository.findCampaigns()
}

func (cs *CampaignService) getCampaign(campaignId string) (Campaign, error) {
	return cs.repository.getCampaignByID(campaignId)
}

func (cs *CampaignService) deleteCampaign(campaignId string) error {
	return cs.repository.deleteCampaign(campaignId)
}

func (cs *CampaignService) createCampaign(r CreateCampaignRequest) (Campaign, error) {
	return cs.repository.createCampaign(r)
}

func (cs *CampaignService) updateCampaign(campaignId string, r UpdateCampaignRequest) (Campaign, error) {
	return cs.repository.updateCampaign(campaignId, r)
}
