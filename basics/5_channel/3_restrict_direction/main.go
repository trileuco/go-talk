package main

import (
	"fmt"
)

func process(ch <-chan int) {
	// ch <- 1 // insert is restricted ...
	fmt.Println(<-ch)
}

func enqueue(ch chan<- int, elem int) {
	// <- ch // extract is restricted ...
	ch <- elem
}

func main() {
	ch := make(chan int, 1)
	enqueue(ch, 1)
	process(ch)
}
