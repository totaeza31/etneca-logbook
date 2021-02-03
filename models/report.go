package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllReport struct {
	Report []Report `json:""`
}

type Report struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Header      string             `json:"header" bson:"header"`
	BoatId      primitive.ObjectID `json:"boatId" bson:"boatId"`
	StartDate   time.Time          `json:"-" bson:"startDate"`
	Start       string             `json:"start"`
	Lastest     string             `json:"lastest"`
	LastestDate time.Time          `json:"-" bson:"lastestDate"`
	BoatName    []boatName         `json:"-"`
	Name        boatname           `json:"boatname"`
}

type DetailReport struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date     time.Time          `json:"date" bson:"date"`
	Detail   string             `json:"detail"`
	Agent    agent              `json:"agent" bson:"agent"`
	ReportId primitive.ObjectID `json:"reportId,omitempty" bson:"reportId,omitempty"`
}

type boatName struct {
	BoatName boatname `json:"boatname" bson:"boatname"`
}
