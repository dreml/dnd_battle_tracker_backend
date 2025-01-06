package characters

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CharactersRepository struct {
	collection *mongo.Collection
}

func NewCharactersRepository(db *mongo.Database) *CharactersRepository {
	return &CharactersRepository{
		collection: db.Collection("characters"),
	}
}

func (cr *CharactersRepository) findCharacters() ([]Character, error) {
	cursor, err := cr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var characters []Character
	if err = cursor.All(context.TODO(), &characters); err != nil {
		return nil, err
	}

	return characters, nil
}

func (cr *CharactersRepository) findCharacterByID(characterId string) (Character, error) {
	id, _ := primitive.ObjectIDFromHex(characterId)
	filter := bson.M{"_id": id}

	var character Character
	err := cr.collection.FindOne(context.Background(), filter).Decode(&character)
	if err != nil {
		return Character{}, err
	}

	return character, nil
}

func (cr *CharactersRepository) createCharacter(r CreateCharacterRequest) (Character, error) {
	character := Character{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		PlayerName:  r.PlayerName,
		Avatar:      r.Avatar,
		Health:      r.Health,
		Armor:       r.Armor,
		CampaignId:  r.CampaignId,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	_, err := cr.collection.InsertOne(context.Background(), character)
	if err != nil {
		return Character{}, err
	}

	return character, nil
}

func (cr *CharactersRepository) updateCharacter(characterId string, r UpdateCharacterRequest) (Character, error) {
	id, _ := primitive.ObjectIDFromHex(characterId)

	var character Character
	filter := bson.M{"_id": id}
	err := cr.collection.FindOne(context.Background(), filter).Decode(&character)
	if err != nil {
		return Character{}, err
	}

	updated := false
	if r.Name != nil {
		character.Name = *r.Name
		updated = true
	}
	if r.PlayerName != nil {
		character.PlayerName = *r.PlayerName
		updated = true
	}
	if r.Avatar != nil {
		character.Avatar = *r.Avatar
		updated = true
	}
	if r.Health != nil {
		character.Health = *r.Health
		updated = true
	}
	if r.Armor != nil {
		character.Armor = *r.Armor
		updated = true
	}
	if r.CampaignId != nil {
		character.CampaignId = *r.CampaignId
		updated = true
	}

	if !updated {
		return character, nil
	}

	character.DateUpdated = time.Now()

	_, err = cr.collection.UpdateByID(context.Background(), id, bson.M{"$set": character})
	if err != nil {
		return Character{}, err
	}

	return character, nil
}

func (cr *CharactersRepository) deleteCharacter(characterId string) error {
	id, _ := primitive.ObjectIDFromHex(characterId)
	filter := bson.M{"_id": id}

	_, err := cr.collection.DeleteOne(context.Background(), filter)

	return err
}
