package main

import (
	"log"

	deimbotapi "github.com/miniapp-io/de-im-bot-go"
)

func main() {
	botToken := "API_TOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	setCommands := deimbotapi.NewSetMyCommands(deimbotapi.BotCommand{
		Command:     "help",
		Description: "a help command",
	})

	res, err := bot.Request(setCommands)
	if err != nil {
		log.Fatal("Unable to set commands")
	}
	log.Printf("set command successfully,result: %s", string(res.Result))

	commands, err := bot.GetMyCommands()
	if err != nil {
		log.Fatal("Unable to get commands")
	}

	log.Printf("get command successfully,result: %v", commands)

	if len(commands) != 1 {
		log.Fatal("Incorrect number of commands returned")
	}

	if commands[0].Command != "help" || commands[0].Description != "a help command" {
		log.Fatal("Commands were incorrectly set")
	}

	setCommands = deimbotapi.NewSetMyCommandsWithScope(deimbotapi.NewBotCommandScopeAllPrivateChats(), deimbotapi.BotCommand{
		Command:     "private",
		Description: "a private command",
	})

	if _, err := bot.Request(setCommands); err != nil {
		log.Fatal("Unable to set commands")
	}

	commands, err = bot.GetMyCommandsWithConfig(deimbotapi.NewGetMyCommandsWithScope(deimbotapi.NewBotCommandScopeAllPrivateChats()))
	if err != nil {
		log.Fatal("Unable to get commands")
	}

	if len(commands) != 1 {
		log.Fatal("Incorrect number of commands returned")
	}

	if commands[0].Command != "private" || commands[0].Description != "a private command" {
		log.Fatal("Commands were incorrectly set")
	}

}
