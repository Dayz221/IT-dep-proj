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

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	bh.HandleMessage(handlers.StartHandler, th.CommandEqual("start"))

	bh.Start()
}
