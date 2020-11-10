package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"paymenthub/models"
	"paymenthub/modules"
)

type AccountRepository struct{}
type AccountBaseRepository interface {
	FetchAll() ([]*models.Account, error)
	FindById(id string) (*models.Account, error)
	StoreAccount(a *models.Account) (bool, error)
	UpdateAccount(id string, a *models.Account) (bool, string, error)
	Delete(id string) (bool, error)
}

var accountCollection = modules.Database("accounts")

func (repo AccountRepository) StoreAccount(newAccount models.Account) (bool, string, error) {
	payment, err := accountCollection.InsertOne(context.TODO(), newAccount)
	if err != nil {
		log.Fatal(err)
		return false, "", err
	}
	return true, payment.InsertedID.(primitive.ObjectID).Hex(), nil
}
