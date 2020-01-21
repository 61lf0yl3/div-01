package main

import (
	"fmt"
	"strings"
)

func main() {

	// result := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// AdvancedSortWordArr(result, strings.Compare)
	// fmt.Println(result)
	// fmt.Println()
	// result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// AdvancedSortWordArr(result, strings.Compare)
	// fmt.Println(result)
	// fmt.Println()
	// result = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// AdvancedSortWordArr(result, strings.Compare)
	// fmt.Println(result)
	// fmt.Println()
	// result = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// AdvancedSortWordArr2(result, strings.Compare)
	// fmt.Println(result)
	// fmt.Println()
	result := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	AdvancedSortWordArr(result, func(a, b string) int { return strings.Compare(b, a) })
	fmt.Println(result)
	// fmt.Println()
	// result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	// AdvancedSortWordArr2(result, strings.Compare)
	// fmt.Println(result)
}

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i1 := range array {
		for i2 := range array {
			if f(array[i1], array[i2]) != 1 {
				array[i1], array[i2] = array[i2], array[i1]
			}
		}
	}
}

// func AdvancedSortWordArr2(array []string, f func(a, b string) int) {
// 	for i1 := range array {
// 		for i2 := range array {
// 			if Compare2(array[i1], array[i2]) == 1 {
// 				array[i1], array[i2] = array[i2], array[i1]
// 			}
// 		}
// 	}
// }

// func Compare(a, b string) int {
// 	if a < b {
// 		return 1
// 	}
// 	return -1
// }

// func Compare2(b, a string) int {
// 	if a < b {
// 		return 1
// 	}
// 	return -1
// }
