package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// producer 1
	ch1 := make(chan int)
	rand.Seed(time.Now().Unix())
	go func() {
		for i := 0; i < 8; i++ {
			ch1 <- rand.Intn(100)
		}
	}()

	// producer 2
	ch2 := make(chan string)
	names := []string{"bob", "alice"}
	go func() {
		for i := 0; i < 2; i++ {
			ch2 <- names[rand.Intn(2)]
		}
	}()
	consuming := true

	// consumer will wait 5 seconds for activity
	for consuming {
		select {
		case number := <-ch1:
			fmt.Println(number)
		case name := <-ch2:
			fmt.Println(name)
		case <-time.After(time.Second * 5):
			consuming = false
		}
		time.Sleep(1 * time.Second)
	}
}
