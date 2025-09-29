package main

import (
	"fmt"
	"strings"
)

func wordRiverse(str string) string {
	arrStr := strings.Fields(str) //разделяем стрингу на чанки
	n := len(arrStr) - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 { //меняем местами слова
		arrStr[i], arrStr[n-i] = arrStr[n-i], arrStr[i]
	}
	var newString string

	for i := range arrStr {
		newString += arrStr[i] + " " //переписываем в формат string
	}
	return newString
}

func main() {
	str := "boy t-shirt black"
	fmt.Println(wordRiverse(str))
}
