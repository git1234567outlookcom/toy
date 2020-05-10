package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Technology struct {
	Id          string             `json:"id,omitempty" bson:"-"`
	ObjectId    primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

func (t *Technology) SetId() *Technology {
	t.Id = t.ObjectId.Hex()
	return t
}
func (t *Technology) SetObjectId() *Technology {
	if t.Id != "" && len(t.Id) == 24 {
		hex, _ := primitive.ObjectIDFromHex(t.Id)
		t.ObjectId = hex
	}
	return t
}
