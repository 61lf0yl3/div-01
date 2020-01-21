package main

import (
	"fmt"
)

func main() {
	result := ConvertBase("4506C", "0123456789ABCDEFM", "choumi")
	fmt.Println(result)
	fmt.Println()
	result = ConvertBase("babcbaaaaabcb", "abac", "0123456789ABCDEF")
	fmt.Println(result)
	fmt.Println()
	result = ConvertBase("0", "0123456789", "01")
	fmt.Println(result)
	fmt.Println()
	result = ConvertBase("u", "choumi", "Zone01")
	fmt.Println(result)
	fmt.Println()
	result = ConvertBase("683241", "0123456789", "0123456789")
	fmt.Println(result)
	fmt.Println()
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nb := AtoiBase(nbr, baseFrom)
	return PrintNbrBase(nb, baseTo)
}

func AtoiBase(s string, base string) int {
	var res int
	for i1 := range s {
		for i2 := range base {
			if s[i1] == base[i2] {
				res = res*len(base) + i2
			}
		}
	}
	//fmt.Println("res from Atoibase", res)
	return res
}

func PrintNbrBase(nbr int, base string) string {
	var res string
	for i1 := nbr; i1 > 0; i1 = i1 / len(base) {

		res = string(base[i1%len(base)]) + res

	}
	return res
}
