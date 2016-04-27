package MiscLib

import (
	"os"

	"github.com/mgutz/ansi"
)

var ColorRed string
var ColorYellow string
var ColorGreen string
var ColorReset string

func init() {
	ColorRed = ""
	ColorYellow = ""
	ColorGreen = ""
	ColorReset = ""
	if !StdErrPiped() { // check if stderr is terminal
		ColorRed = ansi.ColorCode("red:black")
		ColorYellow = ansi.ColorCode("yellow:black")
		ColorGreen = ansi.ColorCode("green:black")
		ColorReset = ansi.ColorCode("reset")
	}
}

func InArray(s string, arr []string) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}

func StdErrPiped() bool {
	fi, _ := os.Stderr.Stat() // get the FileInfo struct describing the standard input.

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true // output is piped to file
	}

	return false
}

func StdOutPiped() bool {
	fi, _ := os.Stdout.Stat() // get the FileInfo struct describing the standard input.

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true // output is piped to file
	}

	return false
}

func StdInPiped() bool {
	fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true // Data from pipe
	}

	return false
}
