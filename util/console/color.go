package console

import (
	"fmt"
)

var (
	RESET  = []byte{27, 91, 48, 109}
	YELLOW = []byte{27, 91, 51, 51, 109}
	RED    = []byte{27, 91, 51, 49, 109}
	CYAN   = []byte{27, 91, 51, 54, 109}
	GREEN  = []byte{27, 91, 51, 50, 109}
)

func Color(color []byte, s string) string {
	return fmt.Sprintf("%s%s%s", color, s, RESET)
}
