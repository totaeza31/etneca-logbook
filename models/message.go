package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Message string "success"
	Result  bool   "true"
}

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   firstName          `bson:"firstname"`
	LastName    lastName           `bson:"lastname"`
	Address     address            `bson:"address"`
	Brithday    time.Time
	Fax         string
	Email       string
	Password    string
	Tel         string
	NamePost    namePost    `bson:"namePost"`
	AddressPost addressPost `bson:"addressPost"`
	Picture     string
}

type firstName struct {
	Th string
	En string
	Bu string
}

type lastName struct {
	Th string
	En string
	Bu string
}
type address struct {
	Th string
	En string
	Bu string
}

type namePost struct {
	Th string
	En string
	Bu string
}
type addressPost struct {
	Th string
	En string
	Bu string
}
