package greets

import (
	"fmt"
)

func SayHello() {
	printGreet("Hello !!")
}

func printGreet(greet string) {
	fmt.Println(greet)
}
