package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllGoods struct {
	Goods []Goods `json:""`
}

type Goods struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code        string             `json:"code" bson:"code"`
	ReceiptCode string             `json:"receiptCode" bson:"receiptCode"`
	Name        string             `json:"name" bson:"name"`
	Cost        int64              `json:"cost" bson:"cost"`
	Remark      string             `json:"remark" bson:"remark"`
}
