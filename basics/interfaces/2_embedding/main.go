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

// cohesive interfaces
type Vehicle interface {
	Engine
	AC
	Tires
}

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
