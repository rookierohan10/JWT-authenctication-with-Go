package databaseConnection

import (
	"context"
	"example/models"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *models.DBConnection{
	MongoDB := os.Getenv("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal("Error crearing Database Connection") 
	}
	fmt.Println("Successfully connected to MongoDB")
	conn := models.DBConnection{
		Conn: client,
	}

	return &conn
}


func OpenCollection(client models.DBConnection, collectionName string) *mongo.Collection{
	collection := client.Conn.Database("cluster0").Collection(collectionName)
	return collection
}
