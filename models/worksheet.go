package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllWorkSheet struct {
	Goods []Goods `json:""`
}

type WorkSheet struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company     string             `json:"company" bson:"company"`
	ReceiptCode string             `json:"receiptCode" bson:"receiptCode"`
	Name        string             `json:"name" bson:"name"`
	Cost        int64              `json:"cost" bson:"cost"`
	Remark      string             `json:"remark" bson:"remark"`
}
