package characters

import (
	"battle_tracker/pkg/common"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	service *CharacterService
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		service: NewCharacterService(db),
	}
}

func (h *Handler) GetCharacters(c echo.Context) error {
	characters, err := h.service.getAllCharacters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var response GetCharactersResponse
	for _, character := range characters {
		response = append(response, createCharacterResponseFromCharacter(character))
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCharacter(c echo.Context) error {
	characterId := c.Param("characterId")

	character, err := h.service.getCharacter(characterId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createCharacterResponseFromCharacter(character))
}

func (h *Handler) CreateCharacter(c echo.Context) error {
	var request CreateCharacterRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	character, err := h.service.createCharacter(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, createCharacterResponseFromCharacter(character))
}

func (h *Handler) UpdateCharacter(c echo.Context) error {
	characterId := c.Param("characterId")

	var request UpdateCharacterRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	character, err := h.service.updateCharacter(characterId, request)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, createCharacterResponseFromCharacter(character))
}

func (h *Handler) DeleteCharacter(c echo.Context) error {
	characterId := c.Param("characterId")

	err := h.service.deleteCharacter(characterId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "Error"})
	}

	return c.JSON(http.StatusOK, nil)
}
