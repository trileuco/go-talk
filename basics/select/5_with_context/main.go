package main

import (
	"context"
	"fmt"
	"math/rand"
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
			}
			time.Sleep(1 * time.Second)
		}
	}()

	<-time.After(3 * time.Second)
	cancel()
}
