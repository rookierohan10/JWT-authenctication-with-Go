package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `json:"first_name"`
	Password      string             `json:"password"`
	Email         string             `json:"email"`
	Phone         string             `json:"phone"`
	Token         string             `json:"token"`
	User_type     string             `json:"user_type"`
	Refresh_token string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
