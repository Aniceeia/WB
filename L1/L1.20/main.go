package main

import (
	"fmt"
	"strings"
)

func wordRiverse(str string) string {
	arrStr := strings.Fields(str)
	n := len(arrStr) - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		arrStr[i], arrStr[n-i] = arrStr[n-i], arrStr[i]
	}
	var newString string

	for i := range arrStr {
		newString += arrStr[i] + " "
	}
	return newString
}

func main() {
	str := "boy t-shirt black"
	fmt.Println(wordRiverse(str))
}
