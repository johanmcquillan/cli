package cli

import (
	"github.com/mgutz/ansi"
)

type ANSI func(string) string

func NewANSI(ansiString string) ANSI {
	return ANSI(ansi.ColorFunc(ansiString))
}

func (a ANSI) Format(s string) string {
	return a(s)
}
