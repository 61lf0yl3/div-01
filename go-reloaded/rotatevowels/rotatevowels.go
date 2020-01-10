package main

import (
	"fmt"
	"os"
)

func Checkvowels(r rune) bool {
	vowels := "AEIOUaeiou"
	for _, l := range vowels {
		if r == l {
			return true
		}
	}
	return false
}

func main() {
	var str string
	for i := range os.Args[1:] {
		str += os.Args[i+1]
		if i < len(os.Args[1:])-1 {
			str += " "
		}
	}
	var res string
	var change string
	for i2 := len(str) - 1; i2 >= 0; i2-- {
		//fmt.Println("str[i2]:", rune(str[i2]))
		if Checkvowels(rune(str[i2])) {
			change += string(str[i2])
		}
	}
	//fmt.Println("str:", str)
	counter := 0
	for _, l1 := range str {
		if Checkvowels(l1) == true {
			//fmt.Println("change:", change)
			l1 = rune(change[counter])
			counter++
		}
		res += string(l1)
	}
	fmt.Println(res)
}
