package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllEmp struct {
	Employee []Employee
}

type Employee struct {
	ID           string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        primitive.ObjectID `json:"title" bson:"title"`
	Firstname    string             `json:"firstname" bson:"firstname"`
	Lastname     string             `json:"lastname" bson:"lastname"`
	Gender       string             `json:"gender" bson:"gender"`
	IdCard       string             `json:"idCard" bson:"idCard"`
	Nation       string             `json:"nation" bson:"nation"`
	Birthday     time.Time          `json:"birthday" bson:"birthday"`
	Address      string             `json:"address" bson:"address"`
	Phone        string             `json:"phone" bson:"phoneNo"`
	StartDate    time.Time          `json:"startDate" bson:"startDate"`
	EndDate      time.Time          `json:"endDate" bson:"endDate"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Position     primitive.ObjectID `json:"position" bson:"position"`
	Company      primitive.ObjectID `json:"company" bson:"company"`
	EmrFirstname string             `json:"emrFirstname" bson:"emrFirstname"`
	EmrLastname  string             `json:"emrLastname" bson:"emrLastname"`
	EmrPhoneNo   string             `json:"emrPhone" bson:"emrPhone"`
	EmrRelate    string             `json:"emrRelate" bson:"emrRelate"`

	Status     bool      `json:"status" bson:"status"`
	EnsureDate time.Time `json:"ensureDate" bson:"ensureDate"`
}

type AllTitle struct {
	Title []Title
}

type Title struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title title              `json:"title" bson:"title"`
}

type title struct {
	TH string `json:"th" bson:"th"`
	EN string `json:"en" bson:"en"`
}

type AllPosition struct {
	Position []Position           `json:"position" bson:"position"`
}

type Position struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Position position           `json:"position" bson:"position"`
}

type position struct {
	TH string `json:"th" bson:"th"`
	EN string `json:"en" bson:"en"`
}
