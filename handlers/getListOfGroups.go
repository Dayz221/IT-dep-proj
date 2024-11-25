package handlers

import (
	"context"
	"itproj/models"
	"itproj/mongodb"
	"log"

	"github.com/mymmrac/telego"
	"go.mongodb.org/mongo-driver/bson"
)

func GetListOfGroups(bot *telego.Bot, query telego.CallbackQuery) []models.Group {
	user, err := models.GetUserById(query.From.ID)
	if err != nil {
		log.Printf("Че та наебнулось к чертям в createTask %s\n", err)
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

func GetListOfGroupsWithAdmin(bot *telego.Bot, query telego.CallbackQuery) []models.Group {
	listOfGroups := GetListOfGroups(bot, query)
	user, err := models.GetUserById(query.From.ID)
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
