package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"paymenthub/models"
	"paymenthub/modules"
	"time"
)


type PaymentRepository struct{}
type PaymentBaseRepository interface {
	GetInquiry(id string, customerId string, merchantId string)
	ChargePayment(a *models.Payment)
	SetSettlement(id string, settlement float64)
	CancelPayment(id string, status bool) (bool, error)
}

var paymentCollection = modules.Database("payments")

func (repo PaymentRepository) GetInquiry(paymentId string, customerId string, merchantId string) models.Payment {
	var result models.Payment
	objectIdCustomer, _ := primitive.ObjectIDFromHex(customerId)
	objID, _ := primitive.ObjectIDFromHex(paymentId)
	filter := bson.M{
		"$and": []bson.M{
			{"_id": objID},
			{"customer_id": objectIdCustomer},
			{"merchant_id": merchantId},
		},
	}
	_ = paymentCollection.FindOne(context.TODO(), filter).Decode(&result)
	return result
}

func (repo PaymentRepository) ChargePayment(newPayment models.Payment) (paymentId string, description string) {
	payment, err := paymentCollection.InsertOne(context.TODO(), newPayment)
	if err != nil {
		log.Fatal(err)
		return "", ""
	}
	return payment.InsertedID.(primitive.ObjectID).Hex(), newPayment.Description
}
func (repo PaymentRepository) SetSettlement(id string) (bool, time.Time) {
	objID, err := primitive.ObjectIDFromHex(id)
	var oldPayment models.Payment
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_ = paymentCollection.FindOne(context.TODO(), filter).Decode(&oldPayment)
	updatedAt := time.Now()
	update := bson.M{
		"$set": bson.M{
			"settle_amount": oldPayment.RequestAmount,
			"updated_at":    updatedAt,
			"status":        1,
		},
	}
	_, err = paymentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false, updatedAt
	}
	return true, updatedAt
}
func (repo PaymentRepository) CancelPayment(id string) (bool, time.Time) {
	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	updatedAt := time.Now()
	update := bson.M{
		"$set": bson.M{
			"status":     2,
			"updated_at": updatedAt,
		},
	}
	_, err = paymentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false, updatedAt
	}
	return true, updatedAt
}