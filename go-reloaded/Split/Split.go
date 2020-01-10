package main

import (
	"fmt"
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
	//charset := "HA"
	//fmt.Println(NbOfWords(str, charset))
}

func NbOfWords(str, charset string) int {
	//fmt.Println(str)
	//fmt.Println("len of str", len(str))
	//fmt.Println("len of charset", len(charset))
	counter := 0
	//fmt.Println(string(str[0 : 0+len(charset)]))
	//fmt.Println(str[0+len(charset) : 0+2*len(charset)])
	for i1 := range str {
		//fmt.Println(string(str[i1 : i1+len(charset)]))
		//fmt.Println(string(str[i1 : i1+4]))

		if str[i1:i1+len(charset)] == charset && str[i1+len(charset):i1+2*len(charset)] != charset {
			//fmt.Println(string(str[i1 : i1+len(charset)]))
			//fmt.Println(i1)
			counter++
			if i1 == len(str)-2*len(charset) {
				return counter - 1
			}
		}
	}
	return counter - 1
}

func len(str string) int {
	couter := 0
	for range str {
		couter++
	}
	return couter
}

func Split(str, charset string) []string {
	str = charset + str + charset + "  "
	res := make([]string, NbOfWords(str, charset))
	counter := 0
	var newstr string
	//fmt.Println(str)
	for i2 :=0; i {
		if str[i2:i2+len(charset)] != charset {
			newstr += string(str[i2])
			//fmt.Println(newstr)
		}
		if str[i2:i2+len(charset)] == charset /*&& str[i2+1] == 32*/ {
			//fmt.Println(string(str[i2]))
			//fmt.Println(i2)
			fmt.Println(newstr)
			res[counter] = newstr
			newstr = ""
			counter++
			if i2 == 23 /*len(str)-2*len(charset)+2*/ {
				return res
			}
		}
	}
	return res
}
