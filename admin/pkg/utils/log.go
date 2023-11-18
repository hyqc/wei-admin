package utils

import (
	"fmt"
	"time"
)

// PrintfLn printf ln
func PrintfLn(f string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s SERVE %s\n", time.Now().Format(time.RFC3339), f), args...)
}
