package main

import (
	"fmt"
)

type Currency float64

const ExchangeEURToDollar = 1.117
const ExchangeDollarToEur = 0.8949

type Dollar Currency

func (d Dollar) Eur() Euro {
	return Euro(d * ExchangeDollarToEur)
}

type Euro Currency

func (e Euro) Dollar() Dollar {
	return Dollar(e * ExchangeEURToDollar)
}

func main() {

	var euros Euro = 13.5
	var dollars Dollar = 13.5

	fmt.Printf("I will give you %v dollars for %v euros", euros.Dollar(), euros)
	fmt.Println()
	fmt.Printf("I will give you %v euros for %v dollars", dollars.Eur(), dollars)
	// fmt.println(euros + dollars)   // compile error
}
