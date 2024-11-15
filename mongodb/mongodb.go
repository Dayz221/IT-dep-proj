package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB() error {
	mongoUri, exists := os.LookupEnv("MONGO_URI")
	if !exists {
		return fmt.Errorf("ошибка при загрузке URI MongoDB")
	}

	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	return nil
}

func GetUserCollection() *mongo.Collection {
	return client.Database("db").Collection("user")
}

func GetGroupCollection() *mongo.Collection {
	return client.Database("db").Collection("group")
}

func GetTaskCollection() *mongo.Collection {
	return client.Database("db").Collection("task")
}
