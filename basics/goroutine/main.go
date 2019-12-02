package main

import (
	"fmt"
	"sync"
	"time"
)

func asynchronousTask(jobs []int, wg *sync.WaitGroup) {
	for j := 0; j < len(jobs); j++ {
		time.Sleep(1 * time.Second)
		println(fmt.Sprintf("Task %v done !", j))
	}
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg.Add(1)
	go asynchronousTask(jobs, wg)
	wg.Wait()
}
