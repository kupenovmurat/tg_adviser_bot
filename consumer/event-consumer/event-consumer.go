package event_consumer

import (
	"log"
	"time"

	"github.com/kupenovmurat/tg_adviser_bot/events"
)

type Consumer struct {
	fetcher   events.Fether
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fether, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err)
			continue
		}

		if len(gotEvents) == 0 {
			log.Println("[INFO] consumer: no events")
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("[ERR] consumer: %s", err)
			continue
		}

		time.Sleep(1 * time.Second)
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("[ERR] consumer: %s", err)
			continue
		}
	}

	return nil
}
