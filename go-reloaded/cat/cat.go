package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func Println(input string) {
	slice := []rune(input)
	for i := range slice {
		z01.PrintRune(slice[i])
	}
}

func main() {
	args := os.Args[1:]
	len := 0
	for range args {
		len++
	}
	if len == 0 {
		io.Copy(os.Stdin, os.Stdout)
	} else {
		for i := 0; i < len; i++ {
			data, error := ioutil.ReadFile(args[i])
			if error != nil {
				err := "cat: " + args[i] + ": No such file or directory"
				Println(err)
				z01.PrintRune(10)

			}
			Println(string(data))
		}
	}
}
