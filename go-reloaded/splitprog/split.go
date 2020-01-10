package main

import (
	"fmt"
)

func len(input string) int {
	lengh := 0
	for range input {
		lengh++
	}
	return lengh
}

func main() {
	str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(Split(str, "|=choumi=|"))
	fmt.Println()
	str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(Split(str, "!==!"))
	fmt.Println()
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(Split(str, "AFJ"))
	fmt.Println()
	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(Split(str, "<<==123==>>"))
}

func Split(str, charset string) []string {
	str += charset
	r1 := []rune(str)
	l1 := len(string(r1))
	r2 := []rune(charset)
	l2 := len(string(r2))
	nbword := 0
	for i := 0; i <= l1-l2; i++ {
		if str[i:i+l2] == charset {
			nbword++
		}
	}
	arr := make([]string, nbword)
	counter := 0
	j := 0
	for i := counter; i <= l1-l2; i++ {
		if str[i:i+l2] == charset {
			arr[j] = str[counter:i]
			counter = i + l2
			j++
		}
	}
	return arr
}
