package main

import (
	"github.com/trileuco/go-talk/basic/interfaces/format"
	"github.com/trileuco/go-talk/basic/interfaces/printer"
)

func main() {
	upperCase := format.UpperCaseFormatter{}
	lowerCase := format.LowerCaseFormatter{}

	formatters := []format.Formatter{
		upperCase,
		lowerCase,
	}
	printer.PrintMessage("Hello !", formatters)
}
