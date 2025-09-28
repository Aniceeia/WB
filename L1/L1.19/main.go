package main

import "fmt"

func reverse(str string) string {
	rune := []rune(str)
	length := len(rune) - 1 //последний элемент

	for i, j := 0, length; i < j; i, j = i+1, j-1 { //итеритуемся от конца и от начала для переворота
		rune[i], rune[length-i] = rune[length-i], rune[i] //последний элемент станет первым и наоборот
	}

	newString := string(rune) //конвертируем обратно, так как строки в го неизменяемый тип данных
	return newString
}

func main() {
	str := "hello"
	fmt.Println(reverse(str))
}
