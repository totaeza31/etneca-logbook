package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MessageAllHuman struct {
	Result   bool     `json:"result"`
	Message  string   `json:"message,omitempty"`
	AllHuman AllHuman `json:"data,omitempty"`
}

type MessageHuman struct {
	Result  bool   `json:"result"`
	Message string `json:"message,omitempty"`
	Human   Human  `json:"data,omitempty"`
}

type AllHuman struct {
	Human []Human `json:"data,omitempty"`
}

type Human struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Own_firstname string             `json:"firstname" bson:"firstname"`
	Own_lastname  string             `json:"lastname" bson:"lastname"`
}
