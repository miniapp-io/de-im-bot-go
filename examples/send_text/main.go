package main

import (
	"log"

	deimbotapi "github.com/miniapp-io/de-im-bot-go"
)

func main() {
	botToken := "APITOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := deimbotapi.NewMessage("CHAT_ID", "Hello, world!")
	res, err := bot.Request(msg)
	if err != nil {
		log.Fatal("Unable to send text message")
	}
	log.Printf("send text message successfully,result: %s", string(res.Result))
}
