package main

type Engine interface {
	Start()
}

type AC interface {
	SetTemperature(float64)
}

type Tires interface {
	Pressure() float64
}

type Altimeter interface {
	Altitude() float64
}

// cohesive interfaces
type Vehicle interface {
	Engine
	AC
	Tires
}

type AirVehicle interface {
	Engine
	Tires
	Altimeter
}

// Car implementation
type Car struct {
	Temperature float64
}

func (c Car) Start() {

}

func (c *Car) SetTemperature(tmp float64) {
	c.Temperature = tmp
}

func (c Car) Pressure() float64 {
	return 23.5
}

// Plane implementation

type Plane struct {
}

func (p Plane) Start() {
}

func (p Plane) Pressure() float64 {
	return 134.5
}

func (p Plane) Altitude() float64 {
	return 2345.67
}
