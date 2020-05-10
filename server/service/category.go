package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"server/database"
)

type CategoryService struct {
	Dao *mongo.Collection
}

func GetCategoryService() *CategoryService {
	return &CategoryService{Dao: database.GetCollection(database.CollectionCategory)}
}
