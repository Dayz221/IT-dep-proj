package utils

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetListOfUsers(groupId string) []models.User {
	group, err := GetGroupInfoByID(groupId)
	if err != nil {
		log.Printf("Что-то пошло по пизде в GetListOfUsers: %s\n", err)
		return []models.User{}
	}

	var users []models.User
	userCollection := mongodb.GetUserCollection()
	for _, id := range group.Users {
		var user models.User
		if err = userCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
			log.Printf("Что-то пошло по пизде в GetListOfUsers: %s\n", err)
			continue
		}
		users = append(users, user)
	}

	return users
}

func GetListOfAdmins(groupId string) []models.User {
	group, err := GetGroupInfoByID(groupId)
	if err != nil {
		log.Printf("Что-то пошло по пизде в GetListOfAdmins: %s\n", err)
		return []models.User{}
	}

	var users []models.User
	userCollection := mongodb.GetUserCollection()
	for _, id := range group.Admins {
		var user models.User
		if err = userCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
			log.Printf("Что-то пошло по пизде в GetListOfAdmins: %s\n", err)
			continue
		}
		users = append(users, user)
	}

	return users
}
