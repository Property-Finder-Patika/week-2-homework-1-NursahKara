package main

import (
	"fmt"
)

func main() {
	width, height := 5., 12.

	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10
	area = area + 10
	area = area * 2
	area = area / 2
	fmt.Printf("area=%g\n", area)

	area -= 10
	area += 10
	area *= 2
	area /= 2
	fmt.Printf("area=%g\n", area)

	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}
