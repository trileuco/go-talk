package conversion

import (
	"strings"
)

type UpperCaseConverter struct{}

func (e UpperCaseConverter) Apply(text string) string {
	return strings.ToUpper(text)
}

type LowerCaseConverter struct{}

func (e LowerCaseConverter) Apply(text string) string {
	return strings.ToLower(text)
}
