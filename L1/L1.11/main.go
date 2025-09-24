package main

import "fmt"

func intersection(arr1 []int, arr2 []int) []int {
	var result []int
	common := make(map[int]bool) //мапа для записи значений первого массива

	for _, v := range arr1 {
		common[v] = true
	}

	for _, v := range arr2 {
		if common[v] { //проверяем наличие значений второго массива в впервом
			result = append(result, v)
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 3, 5, 6, 8, 10}
	arr2 := []int{1, 4, 5, 10, 6, 7}

	fmt.Println(intersection(arr1, arr2))
}
