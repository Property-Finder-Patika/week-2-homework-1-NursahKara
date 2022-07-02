package main

import "fmt"

func main() {
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2),
	)

	fmt.Println(
		1+4-2,
		(1+4)-2,
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2),
	)
}
