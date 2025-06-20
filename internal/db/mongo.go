package db

import (
	"Jsonb/data"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// MongoDB connection setup
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

// Create single document
func CreateMongo(record data.DummyData) error {
	collection := MongoDB.Collection("json_data")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, record)
	return err
}

// Insert multiple documents
func InsertMongo(records []data.DummyData) error {
	var docs []interface{}
	for _, record := range records {
		docs = append(docs, record)
	}
	collection := MongoDB.Collection("json_data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, docs)
	return err
}

// Find single document by 'id'
func FindMongoByID(id int) (data.DummyData, error) {
	collection := MongoDB.Collection("json_data")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result data.DummyData
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	return result, err
}

// Update documents where age >= minAge
func UpdateMongoCity(minAge int, newCity string) error {
	collection := MongoDB.Collection("json_data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"age": bson.M{"$gte": minAge}}
	update := bson.M{"$set": bson.M{"address.city": newCity}}

	_, err := collection.UpdateMany(ctx, filter, update)
	return err
}

// Delete documents where age < maxAge
func DeleteMongoByAge(maxAge int) error {
	collection := MongoDB.Collection("json_data")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"age": bson.M{"$lt": maxAge}}

	_, err := collection.DeleteMany(ctx, filter)
	return err
}
