package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"reward-point/database"
	"reward-point/models"
)

var ctx = context.TODO()
var CustomerCollection = database.Init("customers")

type PointRepository interface {
	GetCustomer(id string) (models.Customer,error)
	GetAllCustomers() ([]models.Customer,error)
	InsertCustomer(customer *models.Customer) error
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id string) error
}


type repository struct {}
func NewPointRepository() PointRepository {
	return &repository{}
}


func (*repository) InsertCustomer(customer *models.Customer) error {
	_, err := CustomerCollection.InsertOne(ctx, customer)
	return err
}

func (*repository) GetCustomer(id string) (models.Customer,error) {
	customer := models.Customer{}
	err := CustomerCollection.FindOne(ctx, bson.D{{"cis_id",id }}).Decode(&customer)
	if err != nil{
		fmt.Println(err)
	}
	return customer, nil
}

func (*repository) GetAllCustomers() ([]models.Customer,error) {
	var customers []models.Customer
	cur, err := CustomerCollection.Find(ctx,bson.D{})
	if err != nil{
		fmt.Println(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var customer models.Customer
		// & character returns the memory address of the following variable.
		err := cur.Decode(&customer) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		customers = append(customers, customer)
	}

	return customers,nil
}
func (*repository) UpdateCustomer(customer *models.Customer) error{
	filter :=bson.D{{"cis_id",customer.CisId}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set",bson.D{{"created_at",customer.CreatedAt},
		{"updated_at",customer.UpdatedAt},
		{"available_points",customer.AvailablePoints}}}}
	updatedCustomer := models.Customer{}
	err := CustomerCollection.FindOneAndUpdate(ctx,filter,update,&returnOpt).Decode(&updatedCustomer)

	if err != nil{
		fmt.Println(err)
	}
	return err
}
func (*repository) DeleteCustomer(id string) error{
	var deletedCustomer models.Customer
	err := CustomerCollection.FindOneAndDelete(ctx, bson.D{{"cis_id",id }}).Decode(&deletedCustomer)
	fmt.Printf("deleted document %v",deletedCustomer)
	return err
}
