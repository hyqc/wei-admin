package utils

import (
	"fmt"
	"testing"
)

func TestGetConfigEnv(t *testing.T) {
	fmt.Println(GetConfigEnv("mode"))
}
