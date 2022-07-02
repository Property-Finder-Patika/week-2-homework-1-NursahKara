package main

import "fmt"

func main() {
	ratio := 1.0 / 10.0

	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f", ratio)
}
