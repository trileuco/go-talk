package main

import (
	"github.com/trileuco/go-talk/basic/interfaces/conversion"
	"github.com/trileuco/go-talk/basic/interfaces/formatter"
	"github.com/trileuco/go-talk/basic/interfaces/printer"
)

func main() {

	upperCaseConverter := conversion.UpperCaseConverter{}
	lowerCaseConverter := conversion.LowerCaseConverter{}
	spacesFormatter := formatter.SpacesFormatter{}

	textModifiers := []printer.TextModifier{
		upperCaseConverter,
		lowerCaseConverter,
		spacesFormatter,
	}
	printer.PrintMessage("Hello !", textModifiers)
}
