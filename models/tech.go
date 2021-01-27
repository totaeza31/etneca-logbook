package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllTech struct {
	Tech []Tech `json:""`
}

type Tech struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company   company            `json:"company" bson:"company"`
	Firstname firstName          `json:"firstname" bson:"firstname"`
	Lastname  lastName           `json:"lastname" bson:"lastname"`
	Nickname  nickname           `json:"nickname" bson:"nickname"`
	Telephone []string          `json:"telephone" bson:"telephone"`
	Address   address            `json:"address" bson:"address"`
	Remark    remark             `json:"remark" bson:"remark"`
}

type nickname struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}


