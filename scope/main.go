package main

import "fmt"

/* global variable declaration */
var x int = 20

func main() {
	/* local variable declaration in main function */
	var x int = 10
	var y int = 20
	var z int

	fmt.Printf("value of x in main() function = %d\n", x)
	z = sum(x, y)
	fmt.Printf("value of z in main() function = %d\n", z)
}

/* function to add two integers */
func sum(a, b int) int {
	fmt.Printf("value of a in sum() function = %d\n", a)
	fmt.Printf("value of b in sum() function = %d\n", b)

	return a + b
}
