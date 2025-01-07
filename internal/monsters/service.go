package monsters

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MonsterService struct {
	monstersCollection *mongo.Collection
	repository         *MonstersRepository
}

func NewMonsterService(db *mongo.Database) *MonsterService {
	return &MonsterService{
		repository: NewMonstersRepository(db),
	}
}

func (cs *MonsterService) getAllMonsters() ([]Monster, error) {
	return cs.repository.findMonsters()
}

func (cs *MonsterService) getMonster(monsterId string) (Monster, error) {
	return cs.repository.findMonsterByID(monsterId)
}

func (cs *MonsterService) deleteMonster(monsterId string) error {
	return cs.repository.deleteMonster(monsterId)
}

func (cs *MonsterService) createMonster(r CreateMonsterRequest) (Monster, error) {
	return cs.repository.createMonster(r)
}

func (cs *MonsterService) updateMonster(monsterId string, r UpdateMonsterRequest) (Monster, error) {
	return cs.repository.updateMonster(monsterId, r)
}
