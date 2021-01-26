package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllHuman struct {
	Human []Human `json:""`
}

type Human struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Own_firstname string             `json:"firstname" bson:"firstname"`
	Own_lastname  string             `json:"lastname" bson:"lastname"`
}
