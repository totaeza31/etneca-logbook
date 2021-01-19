package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message string "success"
	Result  bool   "true"
}

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   firstName          `bson:"firstname"`
	LastName    lastName           `bson:"lastname"`
	Address     address            `bson:"address"`
	Brithday    primitive.DateTime `json:"brithday,omitempty" bson:"brithday,omitempty"`
	Fax         string             `json:"fax,omitempty" bson:"fax,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Tel         string             `json:"telephone,omitempty" bson:"telephone,omitempty"`
	NamePost    namePost           `bson:"namePost"`
	AddressPost addressPost        `bson:"addressPost"`
	Picture     string             `json:"picture,omitempty" bson:"picture,omitempty"`
}

type firstName struct {
	Th string `json:"th,omitempty" bson:"th,omitempty"`
	En string ""
	Bu string ""
}

type lastName struct {
	Th string `json:"th,omitempty" bson:"th,omitempty"`
	En string `json:"en,omitempty" bson:"en,omitempty"`
	Bu string `json:"bu,omitempty" bson:"bu,omitempty"`
}
type address struct {
	Th string `json:"th,omitempty" bson:"th,omitempty"`
	En string `json:"en,omitempty" bson:"en,omitempty"`
	Bu string `json:"bu,omitempty" bson:"bu,omitempty"`
}

type namePost struct {
	Th string `json:"th,omitempty" bson:"th,omitempty"`
	En string `json:"en,omitempty" bson:"en,omitempty"`
	Bu string `json:"bu,omitempty" bson:"bu,omitempty"`
}
type addressPost struct {
	Th string `json:"th,omitempty" bson:"th,omitempty"`
	En string `json:"en,omitempty" bson:"en,omitempty"`
	Bu string `json:"bu,omitempty" bson:"bu,omitempty"`
}
