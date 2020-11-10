package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var uri  = "mongodb://localhost:27017"
 var collection *mongo.Collection
 var ctx = context.TODO()




//MongoDB Connect to DB ...
func Init(collectionName string)  *mongo.Collection{
	// Replace the uri string with your MongoDB deployment's connection string.
	clientOptions := options.Client().ApplyURI(uri)
	var client, err = mongo.Connect(ctx,clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection = client.Database("reward_point_db").Collection(collectionName)
	fmt.Println("Successfully connected and pinged.")
	return collection
}

