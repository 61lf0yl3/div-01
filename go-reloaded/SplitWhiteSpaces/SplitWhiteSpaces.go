package main

import (
	"fmt"
)

func main() {
	str := "The earliest foundations of what 	wou	ld become \ncomp\nuter science predate the invention of the modern digital computer"
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println()
	str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println()
	str = "aiding in computations such as multiplication and division ."
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println()
	str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println()
	str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println()
	str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(SplitWhiteSpaces(str))
}

func NbOfWords(str string) int {
	counter := 0
	for i1 := range str {
		if (str[i1] == 32 || str[i1] == 10 || str[i1] == 9) && (str[i1-1] != 32 && str[i1-1] != 10 && str[i1-1] != 9) {
			counter++
		}
	}
	return counter
}

func SplitWhiteSpaces(str string) []string {
	str += string(rune(32)) + string(rune(9)) + string(rune(10))
	res := make([]string, NbOfWords(str))
	counter := 0
	var newstr string
	for i2 := range str {
		if str[i2] != 32 && str[i2] != 9 && str[i2] != 10 {
			newstr += string(str[i2])
		}
		if (str[i2] != 32 && str[i2] != 9 && str[i2] != 10) && (str[i2+1] == 32 || str[i2+1] == 9 || str[i2+1] == 10) && i2 < len(str) {
			res[counter] = newstr
			newstr = ""
			counter++
		}
	}
	return res
}

func len(str string) int {
	couter := 0
	for range str {
		couter++
	}
	return couter
}
