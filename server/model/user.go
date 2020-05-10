package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// todo 数据验证 https://www.jianshu.com/p/9ef19d5eac72
type User struct {
	Id       string             `json:"id,omitempty" bson:"-"`
	ObjectId primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Age      int                `json:"age" bson:"age,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Phone    string             `json:"phone" bson:"phone,omitempty"`
}

func (u *User) SetObjectId() *User {
	if u.Id != "" && len(u.Id) == 24 {
		u.ObjectId, _ = primitive.ObjectIDFromHex(u.Id)
	}
	return u
}

func (u *User) SetId() *User {
	u.Id = u.ObjectId.Hex()
	return u
}
