package main

import (
	"fmt"
)

func RecursivePower(nb int, power int) int {
	res := 1
	if power < 0 {
		res = 0
	}
	if power > 0 {
		res = nb * RecursivePower(nb, power-1)
	}
	return res
}

func main() {
	// arg1 := 4
	// arg2 := -3
	fmt.Println(RecursivePower(-7, 1))
	fmt.Println(RecursivePower(-8, -7))
	fmt.Println(RecursivePower(4, 8))
	fmt.Println(RecursivePower(1, 3))
	fmt.Println(RecursivePower(-1, 1))
	fmt.Println(RecursivePower(-6, 5))
}
