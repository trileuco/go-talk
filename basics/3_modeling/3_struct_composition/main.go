package main

import (
	"fmt"
)

type Security struct {
	SystemName string
}

func (s Security) OpenDoors() {
	fmt.Println(fmt.Sprintf("System %s : Doors opened", s.SystemName))
}

type Bank struct {
	Name     string
	Security Security
}

func (b Bank) OpenDoors() {
	b.Security.OpenDoors()
}

func main() {

	b := Bank{
		Name:     "NoCashBank",
		Security: Security{SystemName: "NoCashBankSecurity"},
	}
	b.OpenDoors()
}
