package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Authen struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	FirstName    firstName          `json:"firstname"`
	LastName     lastName           `json:"lastname"`
	AccessToken  string             `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string             `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}
