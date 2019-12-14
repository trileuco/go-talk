package main

import (
	"github.com/trileuco/go-talk/basic/interfaces/conversion"
	"github.com/trileuco/go-talk/basic/interfaces/formatter"
	"github.com/trileuco/go-talk/basic/interfaces/printer"
)

func main() {

	upperCase := conversion.UpperCaseConverter{}
	lowerCase := conversion.LowerCaseConverter{}
	spacesFormatter := formatter.SpacesFormatter{}

	converters := []printer.Converter{
		upperCase,
		lowerCase,
		spacesFormatter,
	}
	printer.PrintMessage("Hello !", converters)
}
