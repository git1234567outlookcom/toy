package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/config"
)

var DB *mongo.Database

const (
	CollectionUser       string = "user"
	CollectionRole       string = "role"
	CollectionArticle    string = "article"
	CollectionCategory   string = "category"
	CollectionTechnology string = "technology"
)

func InitClient() {

	uri := "mongodb://" + config.Config.DB.Uri

	clientOptions := options.Client().ApplyURI(uri)
	if config.Config.DB.Username != "" && config.Config.DB.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: config.Config.DB.Username, Password: config.Config.DB.Password})
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("mongo connect err:", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatalln("mongo ping err:", err)
	}
	if config.Config.DB.Name == "" {
		log.Fatalln("config database name is nil")
	}

	DB = client.Database(config.Config.DB.Name)
	log.Println("Connect", uri, "success")
}

func GetCollection(collection string) *mongo.Collection {
	return DB.Collection(collection)
}
