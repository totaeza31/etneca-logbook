package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllWorkSheet struct {
	WorkSheet []WorkSheet `json:""`
}

type WorkSheet struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company string             `json:"company" bson:"company"`
	Status  primitive.ObjectID `json:"status" bson:"status"`
	Time    time.Time          `json:"StartDate" bson:"StartDate"`
}
