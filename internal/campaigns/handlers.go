package campaigns

import (
	"battle_tracker/pkg/common"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	service *CampaignService
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		service: NewCampaignService(db),
	}
}

func (h *Handler) GetCampaigns(c echo.Context) error {
	campaigns, err := h.service.getAllCampaigns()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var response GetCampaignsResponse
	for _, campaign := range campaigns {
		response = append(response, createCampaignResponseFromCampaign(campaign))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")

	campaign, err := h.service.getCampaign(campaignId)

	fmt.Printf("%+v\n", campaign)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createCampaignResponseFromCampaign(campaign))
}

func (h *Handler) CreateCampaign(c echo.Context) error {
	var request CreateCampaignRequest
	c.Bind(&request)

	campaign, err := h.service.createCampaign(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, createCampaignResponseFromCampaign(campaign))
}

func (h *Handler) UpdateCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")

	var request UpdateCampaignRequest
	c.Bind(&request)

	if request.Name == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	campaign, err := h.service.updateCampaign(campaignId, request)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createCampaignResponseFromCampaign(campaign))
}

func (h *Handler) DeleteCampaign(c echo.Context) error {
	campaignId := c.Param("campaignId")

	err := h.service.deleteCampaign(campaignId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, nil)
}
