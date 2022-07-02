package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Defining the variable as the int64 format
	n1 := int64(2)
	n2 := int64(3)
	n3 := int64(4)
	n4 := int64(5)
	//Performing the type conversion for the  number into binary values of them.
	v1 := strconv.FormatInt(n1, 2)
	v2 := strconv.FormatInt(n2, 2)
	v3 := strconv.FormatInt(n3, 2)
	v4 := strconv.FormatInt(n4, 2)
	fmt.Println("The first value after conversion of the integer", v1)
	fmt.Println("The second value after conversion of the integer", v2)
	fmt.Println("The third value after conversion of the integer", v3)
	fmt.Println("The fourth value after conversion of the integer", v4)
}
