package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Payment struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status        int                `json:"status" bson:"status"`
	Description   string             `json:"description" bson:"description"`
	CustomerId    primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	MerchantId    string             `json:"merchant_id" bson:"merchant_id"`
	RequestAmount float64            `json:"request_amount" bson:"request_amount"`
	SettleAmount  float64            `json:"settle_amount" bson:"settle_amount"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}