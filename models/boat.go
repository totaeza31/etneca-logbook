package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllBoats struct {
	Boat []Boat `json:""`
}

type Boat struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BoatName         boatname           `json:"boatName" bson:"boatName"`
	BoatReg          string             `json:"boatReg" bson:"boatReg"`
	Grosstons        float64            `json:"grosstons" bson:"grosstons"`
	TypeBoat         primitive.ObjectID `json:"boatType" bson:"boatType"`
	Agent            agent              `json:"agent" bson:"agent"`
	BoatMan          boatman            `json:"boatman" bson:"boatman"`
	Report           string             `json:"shipTrackingReport" bson:"shipTrackingReport"`
	DeviceNumer      string             `json:"deviceNumber" bson:"deviceNumber"`
	DeviceModel      string             `json:"deviceModel" bson:"deviceModel"`
	EncBoxNumber     string             `json:"encBoxNumber" bson:"encBoxNumber"`
	PlotterNumber    string             `json:"plotterNumber" bson:"plotterNumber"`
	DcNumber         string             `json:"dcNumber" bson:"dcNumber"`
	Username         string             `json:"username" bson:"username"`
	Gateway          primitive.ObjectID `json:"gateway" bson:"gateway"`
	DeviceStatus     primitive.ObjectID `json:"deviceStatus" bson:"deviceStatus"`
	VmsGen           primitive.ObjectID `json:"vmsGen" bson:"vmsGen"`
	BoatBeam         primitive.ObjectID `json:"boatBeamStatus" bson:"boatBeamStatus"`
	FinancialStatus  primitive.ObjectID `json:"financialStatus" bson:"financialStatus"`
	Anniversary_date time.Time          `json:"_" bson:"anniversaryDate"`
	Anniversary      string             `json:"anniversary,omitempty" bson:"anniversary,omitempty"`
	WarrantyExp_date time.Time          `json:"_" bson:"warrantyExpDate"`
	WarrantyExp      string             `json:"warrantyExp,omitempty" bson:"warrantyExp,omitempty"`
	ReportDate_date  time.Time          `json:"_" bson:"reportDatetime"`
	ReportDate       string             `json:"reportDate,omitempty" bson:"reportDate,omitempty"`
	Remark           []string           `json:"remark" bson:"remark"`
}

type boatname struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type agent struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type boatman struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
