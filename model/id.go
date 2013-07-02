package model

import (
	"fmt"
	"os"
)

type Id string

func random_uuid() string {
	f, _ := os.Open("/dev/urandom")
	defer f.Close()

	bytes := make([]byte, 16)
	f.Read(bytes)

	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}

func RandomId() Id {
	return Id(random_uuid())
}
