package monsters

import (
	"battle_tracker/pkg/common"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	service *MonsterService
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		service: NewMonsterService(db),
	}
}

func (h *Handler) GetMonsters(c echo.Context) error {
	monsters, err := h.service.getAllMonsters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var response GetMonstersResponse
	for _, monster := range monsters {
		response = append(response, createMonsterInfoFromMonster(monster))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetMonster(c echo.Context) error {
	monsterId := c.Param("monsterId")

	monster, err := h.service.getMonster(monsterId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createMonsterResponseFromMonster(monster))
}

func (h *Handler) GetMonsterBySlug(c echo.Context) error {
	monsterSlug := c.Param("monsterSlug")

	monster, err := h.service.getMonster(monsterSlug)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createMonsterResponseFromMonster(monster))
}

func (h *Handler) CreateMonster(c echo.Context) error {
	var request CreateMonsterRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	monster, err := h.service.createMonster(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, createMonsterResponseFromMonster(monster))
}

func (h *Handler) UpdateMonster(c echo.Context) error {
	monsterId := c.Param("monsterId")

	var request UpdateMonsterRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	monster, err := h.service.updateMonster(monsterId, request)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createMonsterResponseFromMonster(monster))
}

func (h *Handler) DeleteMonster(c echo.Context) error {
	monsterId := c.Param("monsterId")

	err := h.service.deleteMonster(monsterId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, nil)
}
