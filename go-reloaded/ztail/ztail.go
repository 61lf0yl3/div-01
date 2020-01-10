package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func len(base []string) int {
	lenght := 0
	for range base {
		lenght++
	}
	return lenght
}

func Atoi(str string) int {
	var res int
	for _, l := range str {
		counter := 0
		for i := 0; i <= 9; i++ {
			if l == rune(i+48) {
				res = res*10 + counter
			}
			counter++
		}
	}
	return res
}

func CheckFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckNumber(str string) bool {
	if Atoi(str) > 0 && Atoi(str) < 9223372036854775807 {
		return true
	}
	return false
}

func PrintErr1(args string) {
	err := "tail: cannot open '" + args + "' for reading: No such file or directory"
	fmt.Printf(err)
	fmt.Printf("\n")
	return
}

func PrintErr2(args string) {
	err := "tail: invalid number of bytes: ‘" + args + "’"
	fmt.Printf(err)
	fmt.Printf("\n")
	return
}

func Ztail(filename, nb string) {
	content, _ := os.Open(filename)
	neededBytes := Atoi(nb)
	info, _ := os.Stat(filename)
	size := info.Size()
	var arr []byte
	for i := 0; i <= int(size); i++ {
		arr = append(arr, byte(i))
	}
	content.Read(arr)
	firstByte := int(size) - neededBytes
	if firstByte >= 0 {
		if firstByte < int(size) {
			for i := firstByte; i < int(size); i++ {
				z01.PrintRune(rune(arr[i]))
			}
		}
	}
	if neededBytes >= int(size) {
		for i2 := 0; i2 < int(size); i2++ {
			z01.PrintRune(rune(arr[i2]))
		}
	}
}

func main() {
	if len(os.Args[1:]) == 3 {
		if os.Args[1] == "-c" {
			if CheckNumber(os.Args[2]) {
				if CheckFile(os.Args[3]) {
					Ztail(os.Args[3], os.Args[2])
					return
				} else if !CheckFile(os.Args[3]) {
					PrintErr1(os.Args[3])
					return
				}
			} else if !CheckNumber(os.Args[2]) {
				PrintErr2(os.Args[2])
				return
			}
		}

		if os.Args[2] == "-c" {
			if CheckNumber(os.Args[3]) {
				if CheckFile(os.Args[1]) {
					Ztail(os.Args[1], os.Args[3])
					return
				} else if !CheckFile(os.Args[1]) {
					PrintErr1(os.Args[1])
					return
				}
			} else if !CheckNumber(os.Args[3]) {
				PrintErr2(os.Args[3])
				return
			}
		} else if !CheckNumber(os.Args[2]) {
			PrintErr2(os.Args[2])
			return
		}

	} else if len(os.Args[1:]) == 4 {
		log := "==> " + os.Args[3] + " <=="
		fmt.Printf(log)
		z01.PrintRune(10)
		if os.Args[1] == "-c" {
			if CheckNumber(os.Args[2]) {
				if CheckFile(os.Args[3]) {
					Ztail(os.Args[3], os.Args[2])
				} else if !CheckFile(os.Args[3]) {
					PrintErr2(os.Args[3])
					return
				}
			} else if !CheckNumber(os.Args[2]) {
				PrintErr1(os.Args[2])
				return
			}
		}
		PrintErr1(os.Args[4])
	}
}

// tail -c 20 file.txt    tail -c 20 file
// ./ztail -c 20 file.txt ./ztail -c 20 file

// tail file.txt -c 23    tail file -c 23
// ./ztail file.txt -c 23  ./ztail file -c 23

// tail -c file 15      tail -c file.txt 15
// ./ztail -c file 15     ./ztail -c file.txt 15

// tail -c 11 file.txt file
// ./ztail -c 11 file.txt file

// tail 11 -c file.txt
// ./ztail 11 -c file.txt
