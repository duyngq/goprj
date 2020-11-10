package models

import (
	"time"
)

type Transaction struct{
	TransactionId string	`json:"transaction_id" bson:"transaction_id"`
	TransactionType int 	`json:"transaction_type" bson:"transaction_type"`
	CreatedAt time.Time 	`json:"created_at" bson:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at" bson:"updated_at"`
	Point int 				`json:"point" bson:"point"`
}