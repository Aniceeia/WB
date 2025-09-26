package main

import "fmt"

func binarySearch(arr []int, num int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 4, 7, 8, 12, 22, 43}
	fmt.Println(binarySearch(arr, 43))
}
