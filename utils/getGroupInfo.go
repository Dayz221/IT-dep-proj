package utils

import (
	"context"
	"itproj/models"
	"itproj/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGroupInfoByID(id string) (*models.Group, error) {
	groupId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var group models.Group
	groupCollection := mongodb.GetGroupCollection()
	if err := groupCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: groupId}}).Decode(&group); err != nil {
		return nil, err
	}

	return &group, nil
}
