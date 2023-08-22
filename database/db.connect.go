package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

func DbInstance() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("Error loading environment variables: %w", err)
	}

	mongoDbURL := os.Getenv("MONGO_URL")

	clientOptions := options.Client().ApplyURI(mongoDbURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB")

	return client, nil
}

var Client *mongo.Client = DbInstance()
