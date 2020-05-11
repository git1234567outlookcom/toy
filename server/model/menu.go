package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Path     string             `json:"path" bson:"path,omitempty"`
}

func (m *Menu) SetObjectId() *Menu {
	if m.Id != "" && len(m.Id) == 24 {
		m.ObjectId, _ = primitive.ObjectIDFromHex(m.Id)
	}
	return m
}
func (m *Menu) SetId() *Menu {
	m.Id = m.ObjectId.Hex()
	return m
}
