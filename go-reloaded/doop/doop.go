package main

import (
	"os"

	"github.com/01-edu/z01"
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

func len(base string) int {
	lenght := 0
	for range base {
		lenght++
	}
	return lenght
}
func length(base []string) int {
	lenght := 0
	for range base {
		lenght++
	}
	return lenght
}

func Itoa(nb int) string {
	var str string
	var newstr string
	if nb < 0 {
		nb = nb * -1
		newstr = "-"
		for i := nb; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	} else if nb == 0 {
		newstr = "0"
	} else if nb > 0 {
		for i := nb; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	}
	return newstr
}

func CheckVal(str string) bool {
	for _, l := range str {
		if l < '0' || l > '9' {
			return false
		}
	}
	return true
}

func CheckOperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" || str == "%" {
		return true
	}
	return false
}

func CheckOverFlow(nb, operator string) bool {
	isnegative := 0
	if nb[0] == '-' {
		isnegative = 1
	}
	if isnegative == 1 {
		if len(nb) > 20 {
			return false
		}
		if nb[1:] > "9223372036854775809" {
			return false
		}
	} else {
		if len(nb) > 19 {
			return false
		}
		if nb > "9223372036854775807" {
			return false
		}
	}
	if nb == "9223372036854775807" && (operator == "+" || operator == "*") {
		return false
	}
	return true
}

func main() {
	if length(os.Args[1:]) != 3 {
		z01.PrintRune('0')
		z01.PrintRune(10)
		return
	} else {
		var res int
		val1 := os.Args[1]
		operator := os.Args[2]
		val2 := os.Args[3]
		if !CheckVal(val1) || !CheckVal(val2) || !CheckOperator(operator) || !CheckOverFlow(val1, operator) || !CheckOverFlow(val2, operator) {
			z01.PrintRune('0')
			z01.PrintRune(10)
			return
			// }

		} else {
			nb1 := Atoi(val1)
			nb2 := Atoi(val2)
			// if nb1 > 9223372036854775807 ||
			if operator == "+" {
				res = nb1 + nb2
			} else if operator == "-" {
				res = nb1 - nb2
			} else if operator == "/" {
				if nb2 == 0 {
					err1 := "No division by 0"
					for _, l1 := range err1 {
						z01.PrintRune(l1)
					}
					z01.PrintRune(10)
					return
				} else {
					res = nb1 / nb2
				}
			} else if operator == "*" {
				res = nb1 * nb2
			} else if operator == "%" {
				if nb2 == 0 {
					err2 := "No modulo by 0"
					for _, l2 := range err2 {
						z01.PrintRune(l2)
					}
					z01.PrintRune(10)
					return
				} else {
					res = nb1 % nb2
				}
			}
			if res > 9223372036854775807 || res < -9223372036854775808 {
				z01.PrintRune('0')
				z01.PrintRune(10)
				return
			}
		}
		newres := Itoa(res)
		for _, l3 := range newres {
			z01.PrintRune(l3)
		}
		z01.PrintRune(10)
		return
	}
}
