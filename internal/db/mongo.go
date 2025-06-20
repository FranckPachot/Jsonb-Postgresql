package db

import (
	"Jsonb/data"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoClient = client
	MongoDB = client.Database(os.Getenv("MONGO_DB"))

	fmt.Println("✅ Connected to MongoDB")
	return nil
}

func InsertMongo(records []data.DummyData) error {
	var docs []interface{}
	for _, record := range records {
		docs = append(docs, record)
	}
	collection := MongoDB.Collection("json_data")
	_, err := collection.InsertMany(context.Background(), docs)
	return err
}
