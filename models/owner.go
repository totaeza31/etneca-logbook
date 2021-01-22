package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageAllOwner struct {
	Result   bool     `json:"result"`
	Message  string   `json:"message,omitempty"`
	AllOwner AllOwner `json:"data,omitempty"`
}

type MessageOwner struct {
	Result  bool   `json:"result"`
	Message string `json:"message,omitempty"`
	Owner   Owner  `json:"data,omitempty"`
}

type AllOwner struct {
	Owner []Owner `json:"owner,omitempty"`
}

type Owner struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company       company            `json:"company" bson:"company"`
	Own_firstname own_firstname      `json:"own_firstname" bson:"own_firstname"`
	Own_lastname  own_lastname       `json:"own_lastname" bson:"own_lastname"`
	Id_card       string             `json:"id_card" bson:"id_card"`
	Birthday_date time.Time          `json:"-" bson:"birthday_date,omitempty"`
	Birthday      string             `json:"birthday" bson:"birthday,omitempty"`
	Address       address            `json:"address" bson:"address"`
	Telephone     telephone          `json:"telephone" bson:"telephone"`
	Fax           string             `json:"fax" bson:"fax"`
	Email         string             `json:"email" bson:"email"`
	Username      string             `json:"username" bson:"username"`
	Password      string             `json:"-" bson:"password"`
	Credit        int64              `json:"credit" bson:"credit"`
	IsActive      bool               `json:"isActive" bson:"isActive"`
	Picture       string             `json:"picture" bson:"picture"`
}

type company struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type own_firstname struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type own_lastname struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type address struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type telephone struct {
	Owner   string `json:"owner"`
	Manager string `json:"manager"`
	Captain string `json:"captain"`
	Finance string `json:"finance"`
	Other   string `json:"other"`
}
