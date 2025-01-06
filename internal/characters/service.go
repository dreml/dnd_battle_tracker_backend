package characters

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CharacterService struct {
	charactersCollection *mongo.Collection
	repository           *CharactersRepository
}

func NewCharacterService(db *mongo.Database) *CharacterService {
	return &CharacterService{
		repository: NewCharactersRepository(db),
	}
}

func (cs *CharacterService) getAllCharacters() ([]Character, error) {
	return cs.repository.findCharacters()
}

func (cs *CharacterService) getCharacter(characterId string) (Character, error) {
	return cs.repository.findCharacterByID(characterId)
}

func (cs *CharacterService) deleteCharacter(characterId string) error {
	return cs.repository.deleteCharacter(characterId)
}

func (cs *CharacterService) createCharacter(r CreateCharacterRequest) (Character, error) {
	return cs.repository.createCharacter(r)
}

func (cs *CharacterService) updateCharacter(characterId string, r UpdateCharacterRequest) (Character, error) {
	return cs.repository.updateCharacter(characterId, r)
}
