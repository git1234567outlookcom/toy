package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Menus    string             `json:"menus" bson:"menus,omitempty"`
}
