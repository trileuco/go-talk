package main

import (
	"errors"
	"fmt"
	"log"
)

type Rocket struct {
	fuel         int
	engineStatus bool
	temperature  int
}

func NewRocket(fuel int, temperature int) *Rocket {
	return &Rocket{fuel: fuel, temperature: temperature}
}

func (r Rocket) Start() error {
	if err := r.checkTemperature(); err != nil {
		fmt.Println(err)
	}
	if err := r.initEngine(); err != nil {
		return err
	}
	return nil
}

func (r Rocket) checkTemperature() error {
	const RecommendedTemperature = 100
	if r.temperature > RecommendedTemperature {
		msg := "temperature warning ! %v is higher than the recommended %v !"
		return errors.New(fmt.Sprintf(msg, r.temperature, RecommendedTemperature))
	}
	return nil
}

func (r Rocket) initEngine() error {
	if r.fuel < 1 {
		return errors.New("Cannot start without fuel")
	}
	r.engineStatus = true
	return nil
}

func main() {

	fmt.Println("---> Rocket 1 is starting !!")
	r1 := NewRocket(10, 50)
	if err := r1.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("---> Rocket 1 started.")

	fmt.Println("---> Rocket 2 is starting !!")
	r2 := NewRocket(10, 150)
	if err := r2.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("---> Rocket 2 started !!")

	fmt.Println("---> Rocket 3 is starting !!")
	r3 := NewRocket(0, 10)
	if err := r3.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("---> Rocket 3 started !!")
}
