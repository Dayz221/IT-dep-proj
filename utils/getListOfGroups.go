package utils

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetListOfGroups(userId int64) []models.Group {
	user, err := models.GetUserById(userId)
	if err != nil {
		log.Printf("Че та наебнулось к чертям в GetListOfGroups %s\n", err)
		return []models.Group{}
	}
	purirum := make([]models.Group, 0, 5)
	for _, groupId := range user.Groups {
		var group models.Group
		err = mongodb.GetGroupCollection().FindOne(context.Background(), bson.D{{Key: "_id", Value: groupId}}).Decode(&group)
		if err != nil {
			log.Printf("Проеб в getlistofgroups %s\n ", err)
			continue
		}
		purirum = append(purirum, group)
	}
	return purirum
}

func GetListOfGroupsWithAdmin(userId int64) []models.Group {
	listOfGroups := GetListOfGroups(userId)
	user, err := models.GetUserById(userId)
	if err != nil {
		log.Printf("Проеб в GetListOfGroupsWithAdmin %s\n ", err)
		return []models.Group{}
	}

	groupsWithAdmin := make([]models.Group, 0, 5)

	for _, el := range listOfGroups {
		for _, admin := range el.Admins {
			if admin == user.ID {
				groupsWithAdmin = append(groupsWithAdmin, el)
				break
			}
		}
	}

	return groupsWithAdmin
}
