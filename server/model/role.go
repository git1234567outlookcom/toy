package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string
	Routers  string `json:"routers"` //  role，
}
