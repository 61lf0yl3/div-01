package main

import "fmt"

func len(base string) int {
	lenght := 0
	for range base {
		lenght++
	}
	return lenght
}

func CheckBase(base string) bool {
	var newstr string
	if len(base) < 2 {
		return false
	}
	for _, l1 := range base {
		if l1 == '-' || l1 == '+' {
			return false
		}
		for _, l2 := range newstr {
			if l1 == l2 {
				return false
			}
		}
		newstr += string(l1)
	}
	return true
}

func main() {
	fmt.Println(AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(AtoiBase("0001", "01"))
	fmt.Println(AtoiBase("5", "013+"))
	fmt.Println(AtoiBase("saDt!p!sI", "CHOUMIisDAcat!"))
	fmt.Println(AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(AtoiBase("thisinputshouldnotmatter", "abcd"))
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
	// "bbbbba" bbbbbac

}

func AtoiBase(s string, base string) int {
	var res int
	if !CheckBase(base) {
		return 0
	} else {
		for i1 := range s {
			for i2 := range base {
				if s[i1] == base[i2] {
					res = res*len(base) + i2
				}
			}
		}
	}
	return res
}
