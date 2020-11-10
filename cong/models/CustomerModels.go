package models

import (
	. "time"
)

type Customer struct {
	CisId string				`json:"cis_id" bson:"cis_id" validate:"required"`
	CreatedAt  Time 			`json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt Time 				`json:"updated_at" bson:"updated_at" validate:"required"`
	AvailablePoints []Point   	`json:"available_points" bson:"available_points" `
}
type Point struct {
	Amount int 					`json:"amount" bson:"amount" validate:"required"`
	ExpiredAt Time 				`json:"expired_at" bson:"expired_at" validate:"required"`
	CreatedAt    Time 			`json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt    Time 			`json:"updated_at" bson:"updated_at" validate:"required"`
	TransactionId    string 	`json:"transaction_id" bson:"transaction_id" validate:"required"`
}



