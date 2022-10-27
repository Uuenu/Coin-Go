package eventconsumer

import (
	"log"
	"telegram-coin-go/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Proccessor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Proccessor, batchsize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchsize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize) // got events from fethcer
		if err != nil {
			log.Printf("[ERR] consume: %s", err.Error())
			continue
		}
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}

	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new events: %s", event.Text)

		if err := c.processor.Proccess(event); err != nil { // send event to processor
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}
	return nil
}
