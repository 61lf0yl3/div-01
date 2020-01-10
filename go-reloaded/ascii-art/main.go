package student

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	PrintCombNew(0, n, "")
	z01.PrintRune('\n')
}

func PrintCombNew(start, n int, out string) {

	if n == 0 {
		if out == "9" || out == "89" || out == "789" || out == "6789" || out == "56789" || out == "456789" || out == "3456789" || out == "23456789" || out == "123456789" {
			for _, j := range out {
				z01.PrintRune(j)
			}
		} else {
			for _, j := range out {
				z01.PrintRune(j)
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	for i := start; i <= 9; i++ {
		str := out + string(i+48)
		PrintCombNew(i+1, n-1, str)
	}
}
