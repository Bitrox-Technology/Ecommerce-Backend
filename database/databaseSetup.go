package database

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

func envMONGOURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	return os.Getenv("MONGOURI")
}

func envDBNAME() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	return os.Getenv("DBNAME")
}

func DBSet() *mongo.Client {
	mongoURI := envMONGOURI()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Failed to connect to monogodb!")
		return nil
	}

	fmt.Println("Successfully connected to mongoDB!!!")

	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
