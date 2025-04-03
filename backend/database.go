package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connexion MongoDB :", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB does not respond", err)
	}

	fmt.Println("MongoDB connection successful")
	MongoClient = client
	MongoDB = client.Database("saas_db")
}
