package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	Id            primitive.ObjectID `json:"id" bson:"_id",omitempty`
	NAME          *string            `json:"name" bson:"name" validate:"required,min=2,max=100",omitempty`
	Team_id       string             `json:"team_id" bson:"team_id" ,omitempty`
	Description   *string            `json:"description" bson:"description" ,omitempty`
	Admins        []string           `json:"admins" bson:"admins" ,omitempty`
	Users         []string           `json:"users" bson:"users" validate:"required" ,omitempty`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
}
