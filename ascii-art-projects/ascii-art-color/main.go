package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const height = 8

var colors = map[string]string{
	"black":   "\033[1;30m%s\033[0m",
	"red":     "\033[1;31m%s\033[0m",
	"green":   "\033[1;32m%s\033[0m",
	"yellow":  "\033[1;33m%s\033[0m",
	"blue":    "\033[1;34m%s\033[0m",
	"magenta": "\033[1;35m%s\033[0m",
	"cyan":    "\033[1;36m%s\033[0m",
	"white":   "\033[1;37m%s\033[0m",
	"orange":  "\033[1;38;5;208m%s\033[0m",
}

var words, filename, saveTo, color string
var colored = make(map[int]string)
var flag = "ascii"
var fonts = []string{"standard", "thinkertoy", "shadow"}
var m = make(map[rune][]string)

// detecting whether if it is usual ascii / fs / ouput / color /
func parseArg(args []string) {
	for i1, arg := range args {
		for _, font := range fonts {
			if arg == font && i1 == len(args)-2 && len(args[len(args)-1]) > 9 && args[len(args)-1][:9] == "--output=" {
				filename = arg + ".txt"
				saveTo = args[len(args)-1][9:]
				flag = "output"
			} else if arg == font && i1 == len(args)-2 && len(args[len(args)-1]) > 8 && args[len(args)-1][:8] == "--color=" {
				filename = arg + ".txt"
				tmp := strings.Split(args[len(args)-1], ",")
				color = tmp[0][8:]
				if _, ok := colors[color]; !ok {
					break
				}
				if len(tmp) > 1 {
					for i := 1; i < len(tmp); i++ {
						index, err := strconv.Atoi(tmp[i])
						check(err)
						colored[index] = color
					}
				}
				flag = "color"
			} else if arg == font && i1 == len(args)-1 {
				filename = arg + ".txt"
				flag = "fs"
			}
		}
		if flag == "ascii" {
			filename = "standard.txt"
			words += arg
			if i1 != len(args)-1 {
				words += " "
			}
		}
	}
}

// extracting letters from fonts and putting them into a map
func parseFontFile() {
	data, err := ioutil.ReadFile(filename)
	check(err)
	arrData := strings.Split(string(data), "\n")
	var runeIt = ' '
	var first = 1
	for i2, line := range arrData {
		if i2 >= first && i2 <= first+height {
			m[runeIt] = append(m[runeIt], line)
			if i2 == first+height {
				runeIt++
				first += height + 1
			}
		}
	}
}

func main() {
	var args = os.Args[1:]
	if len(args) == 0 {
		return
	}
	parseArg(args)
	parseFontFile()

	var file *os.File
	var err1 error
	if flag == "output" {
		file, err1 = os.Create(saveTo)
		check(err1)
	}

	for _, set := range strings.Split(words, "\\n") {
		for line := 0; line < height; line++ {
			for index, runa := range set {
				if flag == "ascii" || flag == "fs" {
					printLine(m[runa][line])
				} else if flag == "output" {
					saveLine(m[runa][line], file)
				} else if flag == "color" {
					if _, ok := colored[index]; ok || len(colored) == 0 {
						colorLine(m[runa][line], color)
					} else {
						printLine(m[runa][line])
					}
				}
			}
			if flag == "output" {
				file.WriteString("\n")
			} else {
				fmt.Println()
			}
		}
	}
}

func printLine(mapStr string) {
	for _, symbol := range mapStr {
		fmt.Print(string(symbol))
	}
}

func colorLine(mapStr, color string) {
	for _, symbol := range mapStr {
		fmt.Printf(colors[color], string(symbol))
	}
}

func saveLine(mapStr string, file *os.File) {
	for _, symbol := range mapStr {
		file.WriteString(string(symbol))
	}
}
