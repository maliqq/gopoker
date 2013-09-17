package util

import (
	"fmt"
	"os"
)

// RandomUuid generate random uuid from /dev/urandom
func RandomUuid() string {
	f, _ := os.Open("/dev/urandom")
	defer f.Close()

	bytes := make([]byte, 16)
	f.Read(bytes)

	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}
