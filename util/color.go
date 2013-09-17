package util

import (
	"fmt"
)

// from Scala Console
var (
	Black      = "\033[30m"
	Blackb     = "\033[40m"
	Blink      = "\033[5m"
	Blue       = "\033[34m"
	Blueb      = "\033[44m"
	Bold       = "\033[1m"
	Cyan       = "\033[36m"
	Cyanb      = "\033[46m"
	Green      = "\033[32m"
	Greenb     = "\033[42m"
	Invisible  = "\033[8m"
	Magenta    = "\033[35m"
	Magentab   = "\033[45m"
	Red        = "\033[31m"
	Redb       = "\033[41m"
	Reset      = "\033[0m"
	Reserved   = "\033[7m"
	Underlined = "\033[4m"
	White      = "\033[37m"
	Whiteb     = "\033[47m"
	Yellow     = "\033[33m"
	Yellowb    = "\033[43m"
)

// Color colorify string
func Color(color string, s string) string {
	return fmt.Sprintf("%s%s%s", color, s, Reset)
}

// Colorf colorify std output
func Colorf(color string, format string, args ...interface{}) {
	fmt.Print(Color(color, fmt.Sprintf(format, args...)))
}
