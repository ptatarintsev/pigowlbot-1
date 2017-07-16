package main

import (
	"log"
	"net/http"
	"pigowlbot/token"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.pigowl.com:8443/"+bot.Token, "fullchain.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS(":8443", "fullchain.pem", "privkey.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}