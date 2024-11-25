package main

import (
	"itproj/handlers"
	"itproj/mongodb"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке переменных окружения: %s", err)
	}

	err = mongodb.InitMongoDB()
	if err != nil {
		log.Fatalf("Ошибка при инициализации MongoDB: %s", err)
	}

	botToken, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("Ошибка при загрузке токена бота")
	}

	bot, err := telego.NewBot(botToken)
	if err != nil {
		log.Fatalf("Ошибка при инициализации бота: %s", err)
	}

	me, _ := bot.GetMe()
	log.Println(me.Username)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	bh.HandleMessage(handlers.StartHandler, th.CommandEqual("start"))
	bh.HandleMessage(handlers.FunctionsHandler, th.CommandEqual("functions"))

	bh.HandleCallbackQuery(handlers.CancelHandler, th.CallbackDataEqual("cancel"))
	bh.HandleCallbackQuery(handlers.BackToFunctionsHandler, th.CallbackDataEqual("backToFunctions"))

	bh.HandleCallbackQuery(handlers.CreateGroupHandler, th.CallbackDataPrefix("createGroup"))
	bh.HandleCallbackQuery(handlers.ShowGroupsHandler, th.CallbackDataPrefix("showGroups"))

	bh.HandleCallbackQuery(handlers.TODO_CALLBACK, th.CallbackDataPrefix("createTask"))

	bh.HandleMessage(handlers.MessageHandler, th.AnyMessageWithText())

	bh.Start()
}
