package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	unique := make(map[string]int)

	var result []string

	for _, v := range arr {
		unique[v]++
		if unique[v] > 1 {
			continue
		}
		result = append(result, v)

	}
	fmt.Println(result)
}
