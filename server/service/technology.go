package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"server/database"
)

type TechnologyService struct {
	*mongo.Collection
}

func GetTechnologyService() *TechnologyService {
	return &TechnologyService{database.GetCollection(database.CollectionTechnology)}
}
