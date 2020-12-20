package cli

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type ANSI func(string) string

func NewANSI(ansiString string) ANSI {
	return ANSI(ansi.ColorFunc(ansiString))
}

func (a ANSI) Format(format string, args ...interface{}) string {
	return a(fmt.Sprintf(format, args...))
}
