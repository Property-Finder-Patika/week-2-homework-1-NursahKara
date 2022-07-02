package main

import (
	"fmt"
	"math"
)

func main() {
	//Defining the variable a, b as the integer variables
	var a, b int = 4, 7
	//Type conversion of the numbers after calculation of square root
	var k float64 = 6 + math.Sqrt(float64(a*a+b*b))
	var l uint = uint(k + 9)
	fmt.Println("The value after conversion of the integer", a, b, l)
}
