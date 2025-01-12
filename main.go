package main

import (
	"flag"
	tgClient "go-mod/clients/telegram"
	event_consumer "go-mod/consumer/event-consumer"
	"go-mod/events/telegram"
	"go-mod/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

//7758712642:AAFNbaHCN33KAkH6cG7r-J_wfKv7QUnZX8I

func mustToken() string {
	token := flag.String("tg-bot-token", "", "Token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}
	return *token
}
