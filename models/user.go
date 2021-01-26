package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   firstName          `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName    lastName           `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Address     address            `json:"address,omitempty" bson:"address,omitempty"`
	Birthday    time.Time          `json:"birthday,omitempty" bson:"birthday,omitempty`
	Fax         string             `json:"fax,omitempty" bson:"fax,omitempty"`
	Email       string             `json:"email,omitempty"  bson:"email,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password"`
	Tel         string             `json:"telephone,omitempty" bson:"telephone,omitempty"`
	NamePost    namePost           `json:"namePost,omitempty" bson:"namePost,omitempty"`
	AddressPost addressPost        `json:"addressPost,omitempty" bson:"addressPost,omitempty"`
	Picture     string             `json:"picture,omitempty" bson:"picture,omitempty"  `
}

type firstName struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type lastName struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}


