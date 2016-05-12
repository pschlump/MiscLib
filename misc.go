package MiscLib

import (
	"os"

	"github.com/mgutz/ansi"
)

/*
```go
Color(s, "red")            // red
Color(s, "red+b")          // red bold
Color(s, "red+B")          // red blinking
Color(s, "red+u")          // red underline
Color(s, "red+bh")         // red bold bright
Color(s, "red:white")      // red on white
Color(s, "red+b:white+h")  // red bold on white bright
Color(s, "red+B:white+h")  // red blink on white bright
Color(s, "off")            // turn off ansi codes
```
Colors

* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white
*/

var ColorRed string
var ColorYellow string
var ColorGreen string
var ColorBlue string
var ColorBlack string
var ColorMagenta string
var ColorCyan string
var ColorReset string

func init() {
	ColorRed = ""
	ColorYellow = ""
	ColorGreen = ""
	ColorBlue = ""
	ColorBlack = ""
	ColorMagenta = ""
	ColorCyan = ""
	ColorReset = ""
	if !StdErrPiped() { // check if stderr is terminal
		ColorRed = ansi.ColorCode("red:black")
		ColorYellow = ansi.ColorCode("yellow:black")
		ColorGreen = ansi.ColorCode("green:black")
		ColorBlue = ansi.ColorCode("blue+b:black")
		ColorBlack = ansi.ColorCode("black:black")
		ColorMagenta = ansi.ColorCode("magenta:black")
		ColorCyan = ansi.ColorCode("cyan:black")
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
