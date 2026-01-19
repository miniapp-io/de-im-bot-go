package main

import (
	"log"

	deimbotapi "github.com/miniapp-io/de-im-bot-go"
)

func main() {
	botToken := "API_TOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := deimbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("chat_id:%s", update.Message.Chat.ID)
		}
	}
}
