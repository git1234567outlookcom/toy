package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//文章类别
type Category struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Intro    string             `json:"intro" bson:"intro,omitempty"`
}

func (c *Category) SetObjectId() *Category {
	if c.Id != "" && len(c.Id) == 24 {
		c.ObjectId, _ = primitive.ObjectIDFromHex(c.Id)
	}
	return c
}

func (c *Category) SetId() *Category {
	c.Id = c.ObjectId.Hex()
	return c
}
