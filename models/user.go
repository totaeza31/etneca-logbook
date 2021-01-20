package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   firstName          `json:"firstname"`
	LastName    lastName           `json:"lastname"`
	Address     address            `json:"address"`
	Brithday    time.Time          `json:"brithday"`
	Fax         string             `json:"fax"`
	Email       string             `json:"email"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Tel         string             `json:"telephone,omitempty" bson:"telephone,omitempty"`
	NamePost    namePost           `json:"namePost"`
	AddressPost addressPost        `json:"addressPost"`
	Picture     string             `json:"picture"`
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
type address struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type namePost struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
type addressPost struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
