package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	unique := make(map[string]int)

	var result []string

	for _, v := range arr {
		unique[v]++ //считаем количество уникальных имен
		if unique[v] > 1 {
			continue //пропускаем добавление в результат если значение уже существует в мар
		}
		result = append(result, v)
	}
	fmt.Println(result)
}
