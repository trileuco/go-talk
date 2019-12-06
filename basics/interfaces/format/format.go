package format

import (
	"strings"
)

type Formatter interface {
	Format(msg string) string
}

type UpperCaseFormatter struct{}

func (e UpperCaseFormatter) Format(msg string) string {
	return strings.ToUpper(msg)
}

type LowerCaseFormatter struct{}

func (e LowerCaseFormatter) Format(msg string) string {
	return strings.ToLower(msg)
}
