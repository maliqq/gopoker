package console

import (
	"fmt"
)

var (
	BLACK = "\033[30m"
	BLACK_B = "\033[40m"
	BLINK = "\033[5m"
	BLUE = "\033[34m"
	BLUE_B = "\033[44m"
	BOLD = "\033[1m"
	CYAN = "\033[36m"
	CYAN_B = "\033[46m"
	GREEN = "\033[32m"
	GREEN_B = "\033[42m"
	INVISIBLE = "\033[8m"
	MAGENTA = "\033[35m"
	MAGENTA_B = "\033[45m"
	RED = "\033[31m"
	RED_B = "\033[41m"
	RESET = "\033[0m"
	REVERSED = "\033[7m"
	UNDERLINED = "\033[4m"
	WHITE = "\033[37m"
	WHITE_B = "\033[47m"
	YELLOW = "\033[33m"
	YELLOW_B = "\033[43m"
)

func Color(color string, s string) string {
	return fmt.Sprintf("%s%s%s", color, s, RESET)
}
