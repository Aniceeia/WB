package main

import "fmt"

func ignorRegistr(arr string) string {
	result := []rune(arr) //преобразовываем в слайс рун
	for i, v := range arr {
		if v >= 'A' && v <= 'Z' {
			result[i] = v + ('a' - 'A') //изменяем регистр через таблицу аски
		}
	}
	return string(result)
}

func uniqueSymbols(str string) bool {
	seen := make(map[rune]bool)
	for _, char := range str {
		if seen[char] { //если булевая переменная уже была присвоена, возвращаем false
			return false
		}
		seen[char] = true //каждому символу присваиваем, булевую переменную true
	}
	return true
}

func main() {
	arr := "abcdefAad"
	lower := ignorRegistr(arr)
	fmt.Println(uniqueSymbols(lower))
}
