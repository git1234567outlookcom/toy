package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id         string             `json:"id,omitempty" bson:"-"`
	ObjectId   primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreateTime int64              `json:"createTime" bson:"create_time,omitempty"`
	UpdateTime int64              `json:"updateTime" bson:"update_time,omitempty"`
	Title      string             `json:"title" bson:"title,omitempty"`
	CategoryId string             `json:"categoryId" bson:"category_id,omitempty"`
	UserId     string             `json:"userId" bson:"user_id,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"` //1:草稿，2：发布
	Public     string             `json:"public" bson:"public,omitempty"` //1:公开，2：不公开
}

func (m *Article) SetObjectId() *Article {
	if m.Id != "" && len(m.Id) == 24 {
		m.ObjectId, _ = primitive.ObjectIDFromHex(m.Id)
	}
	return m
}
func (m *Article) SetId() *Article {
	m.Id = m.ObjectId.Hex()
	return m
}
