package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllTech struct {
	Tech []Tech `json:""`
}

type Tech struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company   string             `json:"company" bson:"company"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Telephone []string           `json:"telephone" bson:"telephone"`
	Address string `json:"address" bson:"address"`
	Email   string `json:"email" bson:"email"`
	Remark  string `json:"remark" bson:"remark"`
}
