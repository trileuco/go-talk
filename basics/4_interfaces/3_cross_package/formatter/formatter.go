package formatter

import (
	"strings"
)

type SpacesFormatter struct{}

func (s SpacesFormatter) Apply(text string) string {
	return strings.Join(strings.Split(text, ""), " ")
}

func (s SpacesFormatter) String() string {
	return "SpacesFormatter"
}
