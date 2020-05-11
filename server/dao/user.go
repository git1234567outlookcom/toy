package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"server/database"
	"server/model"
)

type UserDao struct {
	Client *mongo.Collection
}

func GetUserDao() *UserDao {
	return &UserDao{Client: database.GetCollection(database.CollectionUser)}
}

func (d *UserDao) Save(m *model.User) (string, error) {
	one, err := d.Client.InsertOne(context.Background(), m)
	if err != nil {
		// todo 记录日志
		return "", errors.New("新增失败")
	}
	id := one.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (d *UserDao) FindOneByFilter(filter bson.M) (*model.User, error) {

	one := d.Client.FindOne(nil, filter)
	if one.Err() != nil {
		return nil, one.Err()
	}
	user := new(model.User)
	one.Decode(user)
	return user, nil
}

func (d *UserDao) DeleteByFilter(filter bson.M) error {

	_, err := d.Client.DeleteOne(nil, filter, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) FindList(page *model.Page) ([]*model.User, error) {
	option := options.Find()
	option.SetLimit(page.Default().Limit)
	option.SetSkip(page.Limit * (page.PageNum - 1))

	find, err := d.Client.Find(nil, bson.M{}, option)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for find.Next(nil) {
		m := new(model.User)
		err := find.Decode(m)
		if err != nil {
			return nil, err
		}
		users = append(users, m.SetId())
	}

	page.Count, err = d.Client.CountDocuments(nil, bson.M{})
	if err != nil {
		return nil, err
	}
	page.PageCount = int64(math.Ceil(float64(page.Count) / float64(page.Limit)))
	return users, nil
}
