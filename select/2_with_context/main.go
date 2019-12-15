package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// unlimited producer 1
	ch1 := make(chan int)
	rand.Seed(time.Now().Unix())
	go func() {
		for {
			ch1 <- rand.Intn(100)
		}
	}()

	// unlimited producer 2
	ch2 := make(chan string)
	names := []string{"bob", "alice"}
	go func() {
		for {
			ch2 <- names[rand.Intn(2)]
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	// consumer will wait 5 seconds for context cancel
	go func() {
		consuming := true
		for consuming {
			select {
			case number := <-ch1:
				fmt.Println(number)
			case name := <-ch2:
				fmt.Println(name)
			case <-ctx.Done():
				consuming = false
				log.Print("select terminated")
				wg.Done()
			}
			time.Sleep(1 * time.Second)
		}
	}()
	// Will block main execution until timer ends
	<-time.After(3 * time.Second)
	cancel()
	// Will wait to select proper shutdown
	log.Print("Context cancelled, waiting for select termination")
	wg.Wait()
}
