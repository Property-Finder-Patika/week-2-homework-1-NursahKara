package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	fmt.Printf("%02b = %d\n", 0, 0)
	fmt.Printf("%02b = %d\n", 1, 1)

	fmt.Printf("%08b = %d\n", 4, 4)
	fmt.Printf("%08b = %d\n", 8, 8)
	fmt.Printf("%08b = %d\n", 16, 16)
	fmt.Printf("%08b = %d\n", 32, 32)

	i, _ := strconv.ParseInt("00000010", 2, 64)
	fmt.Println(i)

	i, _ = strconv.ParseInt("00010110", 2, 64)
	fmt.Println(i)
}
