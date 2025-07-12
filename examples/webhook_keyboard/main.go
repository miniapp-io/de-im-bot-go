package main

import (
	"log"
	"net/http"

	deimbotapi "github.com/miniapp-io/de-im-bot-go"
)

var numericKeyboard = deimbotapi.NewReplyKeyboard(
	deimbotapi.NewKeyboardButtonRow(
		deimbotapi.NewKeyboardButton("1"),
		deimbotapi.NewKeyboardButton("2"),
		deimbotapi.NewKeyboardButton("3"),
	),
	deimbotapi.NewKeyboardButtonRow(
		deimbotapi.NewKeyboardButton("4"),
		deimbotapi.NewKeyboardButton("5"),
		deimbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	botToken := "API_TOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := deimbotapi.NewWebhook("http://x.com/" + botToken)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("de-im callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:29339", nil)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			log.Println("message is nil")
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			log.Println("message is not command, update:", update)
			continue
		}

		msg := deimbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = deimbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		log.Println("msg send successfully, msg:", msg.Text)
	}
}
