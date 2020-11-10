package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reward-point/database"
	"reward-point/models"
)

var ctx = context.TODO()
var TransactionCollection = database.Init("transactions")

type TransactionRepository interface {
	GetTransaction(id string) (models.Transaction,error)
	UpdateTransaction(trans *models.Transaction) error
	InsertTransaction(trans *models.Transaction) (string,error)
}


type repository struct {}

func NewTransactionRepository() TransactionRepository {
	return &repository{}
}



func (*repository) GetTransaction(id string) (models.Transaction,error) {
	transaction := models.Transaction{}
	err := TransactionCollection.FindOne(ctx, bson.D{{"transaction_id",id }}).Decode(&transaction)
	if err != nil{
		fmt.Println(err)
	}
	return transaction, nil
}

func (*repository) InsertTransaction(trans *models.Transaction) (string , error) {
	res, err := TransactionCollection.InsertOne(ctx, trans)
	return  res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (*repository) UpdateTransaction(trans *models.Transaction) error  {
	filter :=bson.D{{"transaction_id",trans.TransactionId}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set",bson.D{
		{"transaction_type",trans.TransactionType},
		{"created_at",trans.CreatedAt},
		{"updated_at",trans.UpdatedAt},
		{"point",trans.Point}}}}
	updatedTransaction := models.Transaction{}
	err := TransactionCollection.FindOneAndUpdate(ctx,filter,update,&returnOpt).Decode(&updatedTransaction)

	if err != nil{
		fmt.Println(err)
	}
	return err
}

func (*repository) DeleteTransaction(id string) bool{
	err := TransactionCollection.FindOneAndDelete(ctx, bson.D{{"transaction_id",id }})
	if err != nil{
		return false
	}
	return true
}

