package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id",omitempty`
	First_name    *string            `json:"first_name" bson:"firstName" validate:"required,min=2,max=100",omitempty`
	Last_name     *string            `json:"last_name" bson:"lastName" validate:"min=2,max=100",omitempty`
	Password      *string            `json:"Password" bson:"password"  validate:"required,min=6""`
	Email         *string            `json:"email" validate:"email,required" bson:"email"`
	Phone         *string            `json:"phone" validate:"required" bson:"contactNumber" ,omitempty`
	Token         *string            `json:"token" bson:"token" `
	Refresh_token *string            `json:"refresh_token" bson:"refresh_token" `
	User_id       string             `json:"user_id" bson:"user_id" ,omitempty`
	Roles         []string           `json:"Roles" bson:"roles",omitempty`
	Teams         []string           `json:"Teams" bson:"teams",omitempty`
	Age           int                `json:"age" bson:"age",omitempty`
	Address       string             `json:"address" bson:"address",omitempty`
	Location      string             `json:"location" bson:"location",omitempty`
	Title         string             `json:"title" bson:"title",omitempty`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
}
