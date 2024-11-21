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
	//users := mongodb.GetUserCollection()
	user, err := models.GetUserById(query.From.ID)
	if err != nil {
		log.Printf("че та наебнулось к чертям в createTask %s\n", err)
	}
	purirum := make([]models.Group, 0, 5)
	for _, groupId := range user.Groups {
		var group models.Group
		err = mongodb.GetGroupCollection().FindOne(context.Background(), bson.M{"_id": groupId}).Decode(&group)
		if err != nil {
			log.Printf("Проеб в getlistoofgroups %s\n ", err)
		}
		purirum = append(purirum, group)
	}
	return purirum
}

func getGroupsWithCurrentAdmin(bot *telego.Bot, query telego.CallbackQuery) []models.Group {
	user, err := models.GetUserById(query.From.ID)
	if err != nil {
		log.Printf("Обосрались в поиске групп с админом %s", err)
	}
	gr := GetListOfGroups(bot, query)
	res := make([]models.Group, 0, 5)
	for _, j := range gr {
		for _, adminId := range j.Admins {
			if user.ID == adminId {
				res = append(res, j)
				break
			}
		}
	}
	return res
}
