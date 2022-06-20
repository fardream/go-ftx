package main

import (
	"context"
	"log"
	"os/signal"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/fardream/go-ftx"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	outputChan := make(chan *ftx.ChannelResponse[ftx.TickerChannelUpdate])

	wg.Add(1)
	go func() {
		wg.Done()
		defer close(outputChan)
		if err := ftx.SubscribeTicker(ctx, "wss://ftx.us/ws", "BTC/USD", "ticker", outputChan); err != nil {
			log.Fatal(err)
		}
	}()

	for v := range outputChan {
		spew.Dump(v)
	}
}
