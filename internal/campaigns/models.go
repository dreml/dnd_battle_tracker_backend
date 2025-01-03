package campaigns

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campaign struct {
	ID          primitive.ObjectID  `json:"id"`
	Name        string              `json:"name"`
	DateCreated primitive.Timestamp `json:"dateCreated"`
	DateUpdated primitive.Timestamp `json:"dateUpdated"`
}
