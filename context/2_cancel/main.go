package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-time.After(5 * time.Second)
		cancel()
	}()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go asyncTask(ctx, wg, "A")
	go asyncTask(ctx, wg, "B")
	wg.Wait()
}

func asyncTask(ctx context.Context, wg *sync.WaitGroup, name string) {
mainLoop:
	for {
		select {
		case <-ctx.Done():
			log.Printf("Terminating %s", name)
			break mainLoop
		default:
			log.Printf("Processing %s", name)
		}
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
