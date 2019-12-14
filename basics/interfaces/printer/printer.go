package printer

import (
	"fmt"
)

type Converter interface {
	Apply(msg string) string
}

func PrintMessage(text string, converters []Converter) {
	for _, f := range converters {
		fmt.Println(f.Apply(text))
	}
}
