package main

import (
	. "./config"
	. "./tfsClient"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tkanos/gonfig"
	"log"
	"time"
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("time"),
	),
)

func main() {
	configuration := GetConfig()
	err := gonfig.GetConf("./config.json", &configuration)
	bot, err := tgbotapi.NewBotAPI(configuration.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 500)
	updates.Clear()
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /time"
			case "rcutasudd":
				msg.Text = GetIterationWorks(configuration)
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
		//if update.Message.Text == "time" {
		//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetInterationWorks())
		//	bot.Send(msg)
		//}
	}
}