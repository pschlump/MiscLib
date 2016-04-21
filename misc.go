package MiscLib

import (
	"github.com/mgutz/ansi"
)

var ColorRed string
var ColorReset string

func init() {
	ColorRed = ansi.ColorCode("red:black")
	ColorReset = ansi.ColorCode("reset")
}

func InArray(s string, arr []string) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}
