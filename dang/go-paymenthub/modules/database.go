package modules

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database(collection string) *mongo.Collection {
	mongoDBConnectionString := "mongodb://localhost:27017"
	//mongoDBConnectionString := "mongodb://mongodb"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	return client.Database("db_test").Collection(collection)
}
