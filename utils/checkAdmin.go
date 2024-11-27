package utils

import (
	"context"
	"itproj/models"
	"itproj/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAdmin(groupIdHex string, userIdHex string) bool {
	groupId, err := primitive.ObjectIDFromHex(groupIdHex)
	if err != nil {
		return false
	}

	userId, err := primitive.ObjectIDFromHex(userIdHex)
	if err != nil {
		return false
	}

	var group models.Group
	groupCollection := mongodb.GetGroupCollection()
	if err := groupCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: groupId}}).Decode(&group); err != nil {
		return false
	}

	for _, id := range group.Admins {
		if id == userId {
			return true
		}
	}
	return false
}
