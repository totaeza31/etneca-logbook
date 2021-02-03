package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllBoatType struct {
	BoatType []BoatType `json:""`
}

type BoatType struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

type AllBoatBeam struct {
	BoatBeam []BoatBeam `json:""`
}

type BoatBeam struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

type AllBoatDevice struct {
	BoatDevice []BoatDevice `json:""`
}

type BoatDevice struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

type AllBoatFinance struct {
	BoatFinance []BoatFinance `json:""`
}

type BoatFinance struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

type AllBoatGateway struct {
	BoatGateway []BoatGateway `json:""`
}

type BoatGateway struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}
type AllBoatVgm struct {
	BoatVgm []BoatVgm `json:""`
}

type BoatVgm struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}
