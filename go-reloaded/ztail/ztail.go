package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func length(args []string) int {
	len := 0
	for range args {
		len++
	}
	return len
}

//StrLen returns strings length
func StrLen(str string) int {
	len := 0
	for range str {
		len++
	}
	return len
}

//NumCheck checks numbers validity
func NumCheck(str string, c bool) bool {
	if !c {
		for i, j := range str {
			if j == '-' || j == '+' {
				if i != 0 {
					return false
				}
				if StrLen(str) == 1 {
					return false
				}
			} else if j < '0' || j > '9' {
				return false
			}
		}
	} else {
		for i, j := range str {
			if j == '-' {
				if i != 0 && i != 2 {
					return false
				}
				if StrLen(str) == 3 && i == 2 {
					return false
				}
			} else if j == '+' {
				if i != 2 {
					return false
				}
				if StrLen(str) == 3 && i == 2 {
					return false
				}
			} else if j == 'c' {
				if i != 1 {
					return false
				}
			} else if j < '0' || j > '9' {
				return false
			}
		}
	}
	return true
}

//StrToInt64 returns uint64
func StrToInt64(str string) uint64 {
	var res uint64 = 0
	for _, i := range str {
		var num uint64 = 0
		if i >= '0' && i <= '9' {
			for j := '0'; j < i; j++ {
				num++
			}
			res = res*10 + num
		}
	}
	return res
}

//PrintStr prints strings followed by newline
func PrintStr(str string) {
	fmt.Printf(str)
	fmt.Printf("\n")
}

// CheckFile returns true if file exists
func CheckFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// OverflowCheck uint64
func OverflowCheck(str string) bool {
	if StrLen(str) > 20 {
		return false
	} else if StrLen(str) == 20 && str > "18446744073709551615" {
		return false
	}
	return true
}

//GetBytes reads needed amount of bytes from file
func GetBytes(filename string, neededBytes uint64, plus bool) {
	content, err := os.Open(filename)
	if err != nil {
		PrintStr("tail: cannot open '" + filename + "' for reading: Permission denied")
		// fmt.Println(1)
		os.Exit(1)
	}
	info, _ := os.Stat(filename)
	size := info.Size()
	var arr []byte
	var a uint64 = 0
	for i := a; i <= uint64(size); i++ {
		arr = append(arr, ' ')
	}
	content.Read(arr)
	firstByte := uint64(size) - neededBytes
	if plus {
		firstByte = neededBytes - 1
	}
	if neededBytes > uint64(size) && !plus {
		firstByte = 0
	} else if neededBytes > uint64(size) && plus {
		// fmt.Println(2)
		os.Exit(1)
	}
	if firstByte >= 0 {
		for i := firstByte; i < uint64(size); i++ {
			z01.PrintRune(rune(arr[i]))
		}
	}
}

func main() {
	args := os.Args[1:]
	for i, j := range args {
		if j == "-c" && i == length(args)-1 {
			PrintStr("tail: option requires an argument -- 'c'" + "\n" + "Try 'tail --help' for more information.")
			os.Exit(1)
		}
	}
	for i, j := range args {
		if j == "-c" && !NumCheck(args[i+1], false) && i != length(args)-1 || j[0] == '-' && j[1] == 'c' && !NumCheck(j, true) {
			b := j
			if j == "-c" {
				b = args[i+1]
			}
			if StrLen(j) > 2 && j[:2] == "-c" {
				b = j[2:]
			}
			if b[0] == '-' {
				b = b[1:]
			}
			PrintStr("tail: invalid number of bytes: ‘" + b + "‘")
			os.Exit(1)
		}
	}

	for i, j := range args {
		// fmt.Println(j, NumCheck(args[i+1], false), length(args) )
		if (j == "-c" && NumCheck(args[i+1], false) == true && length(args) > 2) || (NumCheck(j, true) && length(args) > 1) {
			var files []string
			var bytes []string
			for b, a := range args {
				if a == "-c" || b != 0 && args[b-1] == "-c" && NumCheck(a, false) || a[0] == '-' && a[1] == 'c' {
					if StrLen(a) > 2 && a[:2] == "-c" {
						a = a[2:]
					} else if a == "-c" {
						continue
					}
					bytes = append(bytes, a)
				} else {
					files = append(files, a)
				}
			}
			plus := false
			var a uint64 = 0
			var b string
			for _, c := range bytes {
				if StrToInt64(c) > a && OverflowCheck(c) {
					a = StrToInt64(c)
					b = c
				} else if !OverflowCheck(c) {
					if c[0] == '-' {
						c = c[1:]
					}
					PrintStr("tail: invalid number of bytes: ‘" + c + "’: Value too large for defined data type")
					os.Exit(2)
				}
			}
			nb := StrToInt64(bytes[length(bytes)-1])
			if bytes[length(bytes)-1][0] == '+' {
				plus = true
			}
			if b[0] == '+' {
				plus = true
			}
			for a, b := range files {
				if CheckFile(b) {
					if length(files) > 1 && (j == "-c" || j != "-c") {
						if a != 0 && ((a != 1 && !CheckFile(files[0])) || a == 1 && CheckFile(files[0]) ) {
							PrintStr("")
						}
						PrintStr("==> " + b + " <==")
					}
					GetBytes(b, nb, plus)
				} else if !CheckFile(b) {
					PrintStr("tail: cannot open '" + b + "' for reading: No such file or directory")
				}
			}
			// fmt.Println(3)
			os.Exit(0)
		} else {
			// fmt.Println(4)
			// os.Exit(0)
			continue
		}
	}
	
}





