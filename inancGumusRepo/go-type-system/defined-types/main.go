package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")

	fmt.Println(h.Hours(), "hours")

	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))
}
