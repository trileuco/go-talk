package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	close(ch) // necessary or deadlock, we need to signal the for loop that consume is ended.

	for j := range ch {
		fmt.Println(j)
	}

}
