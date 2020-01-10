package main

import (
	"fmt"
)

func Check(str string) bool {
	for i, l := range str {
		if l == '-' || l == '+' {
			if i != 0 {
				return false
			}
		} else if l < '0' || l > '9' {
			return false
		}
	}
	return true
}

func Atoi(str string) int {
	var res int
	isnegative := 0
	if str == "" {
		return 0
	}
	if Check(str) == true {
		if str[0] == '-' {
			isnegative = 1
		}
		for _, l := range str {
			counter := 0
			for i := 0; i <= 9; i++ {
				if l == rune(i+48) {
					res = res*10 + counter
				}
				counter++
			}
		}
		if isnegative == 1 {
			res = res * -1
		}
	}
	return res
}

func main() {
	s := ""
	s2 := "-"
	s3 := "--123"
	s4 := "1"
	s5 := "-3"
	s6 := "8292"
	s7 := "9223372036854775807"
	s8 := "-9223372036854775808"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
