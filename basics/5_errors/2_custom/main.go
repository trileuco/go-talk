package main

import (
	"fmt"
	"time"
)

type SystemErr struct {
	timestamp time.Time
	message   string
}

func NewSystemErr(message string) *SystemErr {
	return &SystemErr{message: message, timestamp: time.Now()}
}

func (se SystemErr) Error() string {
	return fmt.Sprintf("Error was '%s' at %v", se.message, se.timestamp)
}

func main() {
	err := NewSystemErr("Demo error !!")
	fmt.Println(err)
}
