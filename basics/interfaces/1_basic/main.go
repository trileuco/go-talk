package main

type Vehicle interface {
	Colour() string
}

type Car struct {
	PaintColour string
}

func (c Car) Colour() string {
	return c.PaintColour
}
