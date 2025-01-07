package monsters

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MonstersRepository struct {
	collection *mongo.Collection
}

func NewMonstersRepository(db *mongo.Database) *MonstersRepository {
	return &MonstersRepository{
		collection: db.Collection("monsters"),
	}
}

func (mr *MonstersRepository) findMonsters() ([]Monster, error) {
	cursor, err := mr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var monsters []Monster
	if err = cursor.All(context.TODO(), &monsters); err != nil {
		return nil, err
	}

	return monsters, nil
}

func (mr *MonstersRepository) findMonsterByID(monsterId string) (Monster, error) {
	id, _ := primitive.ObjectIDFromHex(monsterId)
	filter := bson.M{"_id": id}

	var monster Monster
	err := mr.collection.FindOne(context.Background(), filter).Decode(&monster)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func (mr *MonstersRepository) findMonsterBySlug(monsterSlug string) (Monster, error) {
	filter := bson.M{"index": monsterSlug}

	var monster Monster
	err := mr.collection.FindOne(context.Background(), filter).Decode(&monster)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func (mr *MonstersRepository) createMonster(r CreateMonsterRequest) (Monster, error) {
	monster := Monster{
		ID:          primitive.NewObjectID(),
		Slug:        r.Slug,
		Name:        r.Name,
		Image:       r.Image,
		Health:      r.Health,
		Armor:       r.Armor,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	_, err := mr.collection.InsertOne(context.Background(), monster)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func (mr *MonstersRepository) updateMonster(monsterId string, r UpdateMonsterRequest) (Monster, error) {
	id, _ := primitive.ObjectIDFromHex(monsterId)

	var monster Monster
	filter := bson.M{"_id": id}
	err := mr.collection.FindOne(context.Background(), filter).Decode(&monster)
	if err != nil {
		return Monster{}, err
	}

	updated := false
	if r.Slug != nil {
		monster.Slug = *r.Slug
		updated = true
	}
	if r.Name != nil {
		monster.Name = *r.Name
		updated = true
	}
	if r.Image != nil {
		monster.Image = *r.Image
		updated = true
	}
	if r.Health != nil {
		monster.Health = *r.Health
		updated = true
	}
	if r.Armor != nil {
		monster.Armor = *r.Armor
		updated = true
	}

	if !updated {
		return monster, nil
	}

	monster.DateUpdated = time.Now()

	_, err = mr.collection.UpdateByID(context.Background(), id, bson.M{"$set": monster})
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func (mr *MonstersRepository) deleteMonster(monsterId string) error {
	id, _ := primitive.ObjectIDFromHex(monsterId)
	filter := bson.M{"_id": id}

	_, err := mr.collection.DeleteOne(context.Background(), filter)

	return err
}
