package main

import (
	"fmt"
	"log"
)

func main() {

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"cat", "dog"}
	for k, v := range animals {
		log.Println(k, v)
	}

	points := make(map[string]int)
	points["player1"] = 3
	points["player2"] = 6
	for name, points := range points {
		log.Println(fmt.Sprintf("Player name %s has %v points", name, points))
	}

	running := true
	for running {
		log.Println("Running while true")
		running = false
	}

	for {
		log.Println("This could print infinite times !!")
		break
	}
}
