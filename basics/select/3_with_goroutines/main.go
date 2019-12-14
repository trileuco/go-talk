package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// producer 1
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// producer 2
	ch2 := make(chan string)
	names := []string{"bob", "alice"}
	rand.Seed(time.Now().Unix())
	go func() {
		for i := 0; i < 100; i++ {
			ch2 <- names[rand.Intn(2)]
		}
		close(ch2)
	}()

	// consumer, will wait until channels are empty
	for {
		select {
		case number, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println(number)
		case name, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println(name)
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
