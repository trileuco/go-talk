package printer

import (
	"fmt"
	"github.com/trileuco/go-talk/basic/interfaces/format"
)

func PrintMessage(msg string, formatters []format.Formatter) {
	for _, f := range formatters {
		fmt.Println(f.Format(msg))
	}
}
