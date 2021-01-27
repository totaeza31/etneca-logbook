package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Boats struct {
	Boat []Boat `json:""`
}

type Boat struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BoatName      boatname           `json:"boatname" bson:"boatname"`
	BoatReg       string             `json:"boat_reg" bson:"boat_reg"`
	Grosstons     float64            `json:"grosstons" bson:"grosstons"`
	TypeBoat      typeBoat           `json:"typeBoat" bson:"typeBoat"`
	Agent         agent              `json:"agent" bson:"agent"`
	BoatMan       boatman            `json:"boatman" bson:"boatman"`
	Address       address            `json:"address" bson:"address"`
	Report        string             `json:"ship_tracking_report" bson:"ship_tracking_report"`
	DeviceNumer   string             `json:"device_number" bson:"device_number"`
	DeviceModel   string             `json:"device_model" bson:"device_model"`
	EncBoxNumber  string             `json:"enc_box_number" bson:"enc_box_number"`
	PlotterNumber string             `json:"plotter_number" bson:"plotter_number"`
	DcNumber      string             `json:"dc_number" bson:"dc_number"`
	Username      string             `json:"username" bson:"username"`
	VmsGen        int64              `json:"vms_gen" bson:"vms_gen"`
	BoatBeam      bool               `json:"boatBeam" bson:"boatBeam"`
	Anniversary   time.Time          `json:"anniversary" bson:"anniversary"`
	WarrantyExp   time.Time          `json:"warranty_exp" bson:"warranty_exp"`
	ReportDate    time.Time          `json:"report_	date" bson:"report_	date"`
	Remark        []string           `json:"remark" bson:"remark"`
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

type typeBoat struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}
