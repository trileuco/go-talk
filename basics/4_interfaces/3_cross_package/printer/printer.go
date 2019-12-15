package printer

import (
	"fmt"
)

type TextModifier interface {
	Apply(msg string) string
}

func PrintMessage(text string, textModifiers []TextModifier) {
	for _, f := range textModifiers {
		fmt.Println(f.Apply(text))
	}
}
