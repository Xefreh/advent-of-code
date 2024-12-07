package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read file. err=%v", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	return lines
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
