package main

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/events/telegram"

	eventconsumer "read-adviser-bot/events/telegram/consumer/event-consumer"
	"read-adviser-bot/events/telegram/storage/files"
)


func main() {
	const (
		tgBotHost   = "api.telegram.org"
		storagePath = "files_storage"
		batchSize   = 100
	)
	//token=flags.Get(token)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
	//fetcher=fetcher.New()
	//processor=processor.New()
	//consumer.Start(fetcher, processor)
}
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal()
	}
	return *token
}
