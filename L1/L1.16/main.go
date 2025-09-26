package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2                           // 7 - значение нашего опорного элемента
	arr[pivot], arr[right] = arr[right], arr[pivot] // 7 и 3 меняем местами
	for i := range arr {
		if arr[i] < arr[right] { //arr[right] = 7, все что меньше 7 двигаем влево
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left] //перемещаем опорный элемент в правильную позицию
	quickSort(arr[:left])                         //повторяем для левой части
	quickSort(arr[left+1:])                       //повторяем для правой
	return arr                                    //правая часть - все что больше опрного значения (7, 8, 9), левая (1, 2, 3, 4, 5, 6)
}

func main() {
	arr := []int{6, 4, 9, 2, 7, 5, 8, 1, 3}
	fmt.Println(quickSort(arr))
}
