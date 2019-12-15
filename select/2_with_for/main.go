package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 1)
	ch1 <- 1
	close(ch1)

	ch2 := make(chan string, 1)
	ch2 <- "Alice"
	close(ch2)

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
		time.Sleep(1 * time.Second)
	}
}
