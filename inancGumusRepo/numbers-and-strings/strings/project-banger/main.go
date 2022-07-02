package main

import (
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]

	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)
}
