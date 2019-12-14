package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	fmt.Println("Used parent Area method")
	return r.Base * r.Height
}

type Square struct {
	Rectangle
}

func (s Square) Area() float64 {
	fmt.Println("Used parent Area method")
	return math.Pow(s.Base, 2)
}

func main() {
	square := Square{Rectangle{
		Base:   14.5,
		Height: 14.5,
	}}
	fmt.Println(square.Area())
}
