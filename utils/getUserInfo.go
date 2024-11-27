package utils

import (
	"context"
	"itproj/models"
	"itproj/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByTgId(tgId int64) (*models.User, error) {
	userCollection := mongodb.GetUserCollection()

	var user models.User
	filter := bson.D{{Key: "user_id", Value: tgId}}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(objIdS string) (*models.User, error) {
	objId, err := primitive.ObjectIDFromHex(objIdS)
	if err != nil {
		return nil, err
	}

	var user models.User
	userCollection := mongodb.GetUserCollection()
	filter := bson.D{{Key: "_id", Value: objId}}
	err = userCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
