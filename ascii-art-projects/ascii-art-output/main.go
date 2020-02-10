package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Escape print the error and exit
func Escape(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

// SaveArt prints charset in ascii art
func SaveArt(letters []string, letnum int, f *os.File) {
	for i := 0; i < 8; i++ {
		j := 8
		_, err := f.WriteString(letters[i])
		Escape(err)
		if len(letters) > 8 {
			for k := 1; k < letnum; k++ {
				_, err := f.WriteString(letters[i+j])
				Escape(err)
				if k == letnum-1 {
					_, err := f.WriteString("\n")
					Escape(err)
				}
				j += 8
			}
		} else {
			_, err := f.WriteString("\n")
			Escape(err)
		}
	}
}

func main() {
	str := os.Args[1]
	filename := os.Args[2] + ".txt"
	saveto := os.Args[3][9:]
	f, err := os.Create(saveto)
	Escape(err)
	file, err := os.Open(filename) // open template, save the content
	Escape(err)
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	arr := strings.Split(string(content), "\n") // split file's content by new lines and add to slice
	help := strings.Split(str, "\\n")           // divide our str by new lines
	for _, set := range help {
		var letters []string
		if len(set) == 0 { // if it is plain new lline
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
			SaveArt(letters, letnum, f)
		}
	}
}
