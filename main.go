package main

import (
	. "./config"
	. "./tfs/tfsClient"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func main() {
	configuration := GetConfig()
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
		if update.Message == nil {
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
			_, sendErr := bot.Send(msg)
			if sendErr != nil {
				log.Panic(sendErr)
			}
		}
	}
}