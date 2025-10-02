package main

import "fmt"

func ignorRegistr(arr string) string {
	result := []rune(arr)
	for i, v := range arr {
		if v >= 'A' && v <= 'Z' {
			result[i] = v + ('a' - 'A')
		}
	}
	return string(result)
}

func uniqueSymbols(str string) bool {
	seen := make(map[rune]bool)
	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	arr := "abcdefAad"
	lower := ignorRegistr(arr)
	fmt.Println(uniqueSymbols(lower))
}
