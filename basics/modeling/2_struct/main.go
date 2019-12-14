package main

import (
	"fmt"
)

type Dog struct {
	Name, ChipId string
	Owner        string
	Age          int
}

func (d Dog) Bark() {
	fmt.Println(fmt.Sprintf("%s: Wouf !!", d.Name))
}

func makeItBark(d Dog) {
	d.Bark()
}

func main() {
	dog := Dog{
		Name:  "Pakis",
		Owner: "Eloy",
		Age:   11,
	}
	makeItBark(dog)

	fmt.Printf("Created dog: %v", dog) // beware of zero value "" (ChipId)
}
