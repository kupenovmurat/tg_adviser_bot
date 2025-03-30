package main

import (
	"flag"
	"log"

	tgClient "github.com/kupenovmurat/tg_adviser_bot/clients/telegram"
	event_consumer "github.com/kupenovmurat/tg_adviser_bot/consumer/event-consumer"
	"github.com/kupenovmurat/tg_adviser_bot/events/telegram"
	"github.com/kupenovmurat/tg_adviser_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Println("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}

}

func mustToken() string {
	token := flag.String("token-bot-api", "", "token for telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not set")
	}

	return *token
}
