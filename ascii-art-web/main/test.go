package main

import (
	"fmt"

	ascii ".."
)

func main() {
	arr := []string{"mam", "shadow"}
	out, err := ascii.Asciify(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
