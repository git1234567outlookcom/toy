package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Base struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
}

func (m *Base) SetObjectId() *Base {
	if m.Id != "" && len(m.Id) == 24 {
		m.ObjectId, _ = primitive.ObjectIDFromHex(m.Id)
	}
	return m
}
func (m *Base) SetId() *Base {
	m.Id = m.ObjectId.Hex()
	return m
}

type ResErr struct {
	//Error     interface{} `json:"error,omitempty"` //  测试使用
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

type Res struct {
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	PageInfo  interface{} `json:"pageInfo,omitempty"`
}
