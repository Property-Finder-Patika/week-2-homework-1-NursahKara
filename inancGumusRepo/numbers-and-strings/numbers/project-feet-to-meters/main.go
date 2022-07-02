package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, _ := strconv.ParseFloat(os.Args[1], 64)
	f := c*1.8 + 32

	fmt.Printf("%g ºC is %g ºF\n", c, f)

	fmt.Printf("%g ºF\n", f)
}
