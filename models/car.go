package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents the MongoDB user document
type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Model   int                `json:"age,omitempty" bson:"age,omitempty"`
}
