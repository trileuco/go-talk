package main

import (
	"fmt"
	"sync"
	"time"
)

func asynchronousTask(ch chan int, wg *sync.WaitGroup) {
	for j := range ch {
		time.Sleep(1 * time.Second)
		println(fmt.Sprintf("Task %v done !", j))
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go asynchronousTask(ch, wg)

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, j := range jobs {
		ch <- j
	}
	close(ch)
	wg.Wait()
}
