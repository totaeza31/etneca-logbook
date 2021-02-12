package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllEmp struct {
	Employee []Employee
}

type Employee struct {
	ID            string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Title         primitive.ObjectID `json:"title" bson:"title"`
	Firstname     string             `json:"firstname" bson:"firstname"`
	Lastname      string             `json:"lastname" bson:"lastname"`
	Gender        primitive.ObjectID `json:"gender" bson:"gender"`
	IdCard        string             `json:"idCard" bson:"idCard"`
	Nation        string             `json:"nation" bson:"nation"`
	Birthday      string             `json:"birthday" bson:"-"`
	BirthdayTime  time.Time          `json:"-" bson:"birthday"`
	Address       string             `json:"address" bson:"address"`
	Phone         string             `json:"phoneNo" bson:"phoneNo"`
	StartDate     string             `json:"startDate" bson:"-"`
	StartDateTime time.Time          `json:"-" bson:"startDate"`
	EndDate       string             `json:"endDate" bson:"-"`
	EndDateTime   time.Time          `json:"-" bson:"endDate"`
	Email         string             `json:"email" bson:"email"`
	Password      string             `json:"password" bson:"password"`
	Position      primitive.ObjectID `json:"position" bson:"position"`
	Company       primitive.ObjectID `json:"company" bson:"company"`
	EmrTitle      primitive.ObjectID `json:"emrTitle" bson:"emrTitle"`
	EmrFirstname  string             `json:"emrFirstname" bson:"emrFirstname"`
	EmrLastname   string             `json:"emrLastname" bson:"emrLastname"`
	EmrPhoneNo    string             `json:"emrPhone" bson:"emrPhone"`
	EmrRelate     string             `json:"emrRelate" bson:"emrRelate"`

	Status         bool      `json:"status" bson:"status"`
	EnsureDate     string    `json:"ensureDate" bson:"-"`
	EnsureDateTime time.Time `json:"-" bson:"ensureDate"`
}

type AllTitle struct {
	Title []Title
}
type AllGender struct {
	Gender []Gender
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
	Position []Position `json:"position" bson:"position"`
}

type Position struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Position position           `json:"position" bson:"position"`
}

type position struct {
	TH string `json:"th" bson:"th"`
	EN string `json:"en" bson:"en"`
}

type Gender struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Gender title              `json:"gender" bson:"gender"`
}

type GetEmployee struct {
	ID           string     `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        []Title    `json:"-" bson:"title"`
	TitleName    Title      `json:"titleDetail" bson:"titleDetail"`
	Firstname    string     `json:"firstname" bson:"firstname"`
	Lastname     string     `json:"lastname" bson:"lastname"`
	Gender       []Gender   `json:"-" bson:"gender"`
	Gd           Gender     `json:"gdDetail" bson:"gdDetail"`
	IdCard       string     `json:"idCard" bson:"idCard"`
	Nation       string     `json:"nation" bson:"nation"`
	Birthday     time.Time  `json:"birthday" bson:"birthday"`
	Address      string     `json:"address" bson:"address"`
	Phone        string     `json:"phone" bson:"phoneNo"`
	StartDate    time.Time  `json:"startDate" bson:"startDate"`
	EndDate      time.Time  `json:"endDate" bson:"endDate"`
	Email        string     `json:"email" bson:"email"`
	Password     string     `json:"password" bson:"password"`
	Position     []Position `json:"-"  bson:"position"`
	Pst          Position   `json:"pstDetail" bson:"pstDetail"`
	Company      []Company  `json:"-" bson:"company"`
	Com          Company    `json:"compDetail" bson:"compDetail"`
	EmrTiltle    string     `json:"emrFirstname" bson:"emrFirstname"`
	EmrFirstname []Title    `json:"-" bson:"title"`
	EmrLastname  string     `json:"emrLastname" bson:"emrLastname"`
	EmrPhoneNo   string     `json:"emrPhone" bson:"emrPhone"`
	EmrRelate    string     `json:"emrRelate" bson:"emrRelate"`

	Status     bool      `json:"status" bson:"status"`
	EnsureDate time.Time `json:"ensureDate" bson:"ensureDate"`
}

type AllGetEmployee struct {
	GetEmployee []GetEmployee
}
