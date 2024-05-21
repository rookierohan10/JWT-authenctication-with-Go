package models

import "go.mongodb.org/mongo-driver/mongo"

type DBConnection struct{
	Conn *mongo.Client
}