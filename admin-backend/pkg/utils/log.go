package utils

import (
	"fmt"
	"time"
)

// PrintfLn printf ln
func PrintfLn(f string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf(fmt.Sprintf("%v SERVE %v\n", time.Now().Format(time.RFC3339), f))
		return
	}
	fmt.Printf(fmt.Sprintf("%v SERVE %v \n", time.Now().Format(time.RFC3339), f), args...)
	fmt.Println()
}
