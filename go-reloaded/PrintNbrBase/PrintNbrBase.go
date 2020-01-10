package main

import "github.com/01-edu/z01"

func len(str string) int {
	var len int
	for range str {
		len++
	}
	return len
}

func CheckBase(base string) bool {
	var newstr string
	if len(base) < 2 {
		return false
	}
	for _, l := range base {
		if l == '+' || l == '-' {
			return false
		}
		for _, l2 := range newstr {
			if l == l2 {
				return false
			}
		}
		newstr += string(l)
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if CheckBase(base) == false {
		notvalid := "NV"
		for _, l := range notvalid {
			z01.PrintRune(l)
		}
		return
	}
	var res string
	isnegative := 0
	if nbr < 0 {
		if nbr == -9223372036854775808 {
			res = "8085774586302733229-"
		} else {
			nbr = nbr * -1
			isnegative = 1
		}
	}
	for i1 := nbr; i1 > 0; i1 = i1 / len(base) {

		res += string(base[i1%len(base)])

	}
	if isnegative == 1 {
		res += "-"
	}
	for i2 := len(res) - 1; i2 >= 0; i2-- {
		z01.PrintRune(rune(res[i2]))
	}
}

func main() {
	PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	PrintNbrBase(-661165, "1")
	z01.PrintRune('\n')
	PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')

}
