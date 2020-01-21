package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var args string
	for i := range os.Args[1:] {
		args += os.Args[i+1]
		if i < len(os.Args[1:])-1 {
			args += " "
		}
	}
	var m = make(map[rune]string)
	data, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	var res string
	var counter int
	runeit := ' '
	for _, l := range data {
		if counter > 0 && counter < 8 {
			res += string(l)
		}
		if counter == 8 && l != 10 {
			res += string(l)
		}
		if l == 10 { // if rune = newline
			counter++
		}
		if counter == 9 { // number of newlines
			m[runeit] = res
			runeit++
			counter = 0
			res = ""
		}
	}
	fmt.Println(m)
}

// 	for _, l2 := range strings.Split(args, "\\n") {
// 		for i1 := 0; i1 < 8; i1++ {
// 			for _, l3 := range l2 {
// 				PrintLine(m[l3], i1)
// 			}
// 			z01.PrintRune(10)
// 		}
// 		//z01.PrintRune(10)
// 	}
// }
// func PrintLine(maps string, line int) {
// 	counter := 0
// 	for _, l3 := range maps {
// 		if l3 == '\n' {
// 			counter++
// 		}
// 		if counter > line {
// 			return
// 		}
// 		if counter == line {
// 			if l3 != '\n' {
// 				z01.PrintRune(l3)
// 			}
// 		}
// 	}
// }
