package main

import (
	"fmt"
)

func main() {
	nbits := ActiveBits(15)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(17)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(4)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(11)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(9)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(12)
	fmt.Println(nbits)
	fmt.Println()
	nbits = ActiveBits(2)
	fmt.Println(nbits)
}

func ActiveBits(n int) uint {
	var counter uint
	for n > 0 {
		if n%2 == 1 {
			counter++
		}
		n = n / 2
	}
	return counter
}
