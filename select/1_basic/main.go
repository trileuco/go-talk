package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int, 1)
	ch1 <- 1
	close(ch1)

	ch2 := make(chan string, 1)
	ch2 <- "Alice"
	close(ch2)

	select {
	case number := <-ch1:
		fmt.Println(number)
	case name := <-ch2:
		fmt.Println(name)
	}
	// will print randomly a name or a number
}
