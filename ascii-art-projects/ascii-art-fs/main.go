package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// PrintArt prints charset in ascii art
func PrintArt(letters []string, letnum int) {
	for i := 0; i < 8; i++ {
		j := 8
		fmt.Print(letters[i])
		if len(letters) > 8 {
			for k := 1; k < letnum; k++ {
				fmt.Print(letters[i+j])
				if k == letnum-1 {
					fmt.Print("\n")
				}
				j += 8
			}
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	if len(os.Args) == 3 {
		str := os.Args[1]
		filename := os.Args[2] + ".txt"
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer file.Close()
		content, _ := ioutil.ReadAll(file)
		arr := strings.Split(string(content), "\n") // split file's content by new lines and add to slice
		help := strings.Split(str, "\\n")           // divide our str by new lines
		for _, set := range help {
			var letters []string
			if len(set) == 0 { // if it is plain new line
				fmt.Print("\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n")
			} else {
				var letnum = 0
				for _, letter := range set {
					letnum++
					b := 1
					for i := 0; i < int(letter)-32; i++ { // count first line for given rune to take from content
						b += 9
					}
					for _, line := range arr[b : b+8] {
						letters = append(letters, line)
					}
				}
				PrintArt(letters, letnum)
			}
		}
	}
}
	