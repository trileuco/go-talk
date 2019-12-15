package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Event struct {
	UUID      string    `json:"uuid"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

func main() {
	ev := Event{
		UUID:      "A1234",
		Timestamp: time.Now(),
		Type:      "demo.running",
	}
	data, err := json.Marshal(ev)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
