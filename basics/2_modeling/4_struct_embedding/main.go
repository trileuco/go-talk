package main

import (
	"fmt"
)

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

type Square struct {
	Rectangle
}

func main() {
	square := Square{Rectangle{
		Base:   14.5,
		Height: 14.5,
	}}
	fmt.Println(square.Area())

	fmt.Println(square.Base)
	fmt.Println(square.Height)

	fmt.Println(square.Rectangle.Height)
	fmt.Println(square.Rectangle.Base)
}
