package main

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

func main() {
	const (
		tgBotHost = "api.telegram.org"
	)
	//token=flags.Get(token)
	tgClient:=telegram.New(mustToken())
	//fetcher=fetcher.New()
	//processor=processor.New()
	//consumer.Start(fetcher, processor)
}
func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token=="" {
		log.Fatal()
	}
	return *token
}
 