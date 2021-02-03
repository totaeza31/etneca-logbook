package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllWorkSheet struct {
	WorkSheet []WorkSheet `json:""`
}

type WorkSheet struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company      string             `json:"company" bson:"company"`
	DeviceNumber string             `json:"deviceNumber" bson:"deviceNumber"`
	Status       primitive.ObjectID `json:"status" bson:"status"`
	Time         time.Time          `json:"StartDate" bson:"StartDate"`
}

type WorkSheetRespond struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company      string             `json:"company" bson:"company"`
	DeviceNumber string             `json:"deviceNumber" bson:"deviceNumber"`
	Status       primitive.ObjectID `json:"status" bson:"status"`
	Time         time.Time          `json:"StartDate" bson:"StartDate"`
	BoatDevice   []boatDetail       `json:"-" bson:"boatDetail"`
	BoatName     boatname           `json:"boatName" bson:"boatName"`
	TechDetail   []techDetail       `json:"-" bson:"techDetail"`
	Telephone    []string           `json:"telephone" bson:"telephone"`
}

type boatDetail struct {
	Name boatname `json:"boatName" bson:"boatName"`
}

type techDetail struct {
	Telephone []string `json:"telephone" bson:"telephone"`
}
