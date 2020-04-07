package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const height = 8

func escapeIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	fonts := []string{"standard", "thinkertoy", "shadow"}
	args := os.Args[1:]
	var words, filename, saveTo string
	var ascii, output, fs = false, false, false
	for i, arg := range args {
		for ind, font := range fonts {
			if (arg == font && i == len(args)-2 && args[len(args)-1][:9] == "--output=") || (arg == font && i == len(args)-1) {
				filename = arg + ".txt"
				if i == len(args)-2 {
					saveTo = args[len(args)-1][9:]
					output = true
				} else {
					fs = true
				}
			} else if i == len(args)-1 && ind == len(fonts)-1 && !fs && !output {
				words += arg
				filename = "standard.txt"
				ascii = true
			}
		}
		if !ascii && !output && !fs {
			words += arg
			if i != len(args)-1 {
				words += " "
			}
		}
	}
	data, _ := ioutil.ReadFile(filename)
	arrData := strings.Split(string(data), "\n")
	var m = make(map[rune][]string)
	var runeIt = ' '
	var first = 1
	for index, line := range arrData {
		if index >= first && index <= first+height {
			m[runeIt] = append(m[runeIt], line)
			if index == first+height {
				runeIt++
				first += height + 1
			}
		}
	}
	var file *os.File
	var err error
	if output {
		file, err = os.Create(saveTo)
		escapeIf(err)
	}
	for _, set := range strings.Split(words, "\\n") {
		for line := 0; line < height; line++ {
			for _, runa := range set {
				if ascii || fs {
					printLine(m[runa][line], line)
				} else if output {
					saveArt(m[runa][line], line, file)
				}
			}
			if ascii || fs {
				fmt.Print("\n")
			} else if output {
				file.WriteString("\n")
			}
		}
	}
}

func printLine(mapStr string, line int) {
	for _, symbol := range mapStr {
		fmt.Print(string(symbol))
	}
}

func saveArt(mapStr string, line int, file *os.File) {
	for _, symbol := range mapStr {
		file.WriteString(string(symbol))
	}
}
