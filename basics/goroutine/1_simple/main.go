package main

import (
	"log"
	"time"
)

func jobB() {
	for {
		log.Print("Hello, i will process job B here !!")
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go func() {
		for {
			log.Print("Hello, i will process job A here !!")
			time.Sleep(1 * time.Second)
		}
	}()
	go jobB()
	time.Sleep(10 * time.Second)
}
