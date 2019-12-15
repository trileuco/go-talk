package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	// note that contexts are stackable
	ctx := context.WithValue(context.Background(), "type", "data")
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
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
			log.Printf("Terminating %s with type: %s", name, ctx.Value("type"))
			break mainLoop
		default:
			log.Printf("Processing %s", name)
		}
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
