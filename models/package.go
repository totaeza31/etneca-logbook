package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	Package []Package 
}

type Package struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       name               `json:"name"`
	Desription description        `json:"description,omitempty" bson:"description,omitempty"`
	Remark     remark             `json:"remark,omitempty" bson:"remark,omitempty"`
	Period     []period           `json:"period,omitempty" bson:"period,omitempty`
}

type name struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type description struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type remark struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

type period struct {
	Th Th `json:"th"`
	En En `json:"en"`
	Bu Bu `json:"bu"`
}

type Th struct {
	Month string `json:"month"`
	Cost  string `json:"cost"`
}

type En struct {
	Month string `json:"month"`
	Cost  string `json:"cost"`
}

type Bu struct {
	Month string `json:"month"`
	Cost  string `json:"cost"`
}
