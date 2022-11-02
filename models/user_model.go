package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id",omitempty`
	Name     string             `json:"Name" bson:"firstName",omitempty`
	Roles    []string           `json:"Roles" bson:"roles",omitempty`
	Teams    []string           `json:"Teams" bson:"teams",omitempty`
	Age      int                `json:"age" bson:"age",omitempty`
	Address  string             `json:"address" bson:"address",omitempty`
	Location string             `json:"location" validate:"required",omitempty`
	Title    string             `json:"title" validate:"required",omitempty`
}
