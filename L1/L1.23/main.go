package main

import "fmt"

func deleteElem(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		return arr
	}
	result := make([]int, 0, len(arr)-1)  //создаем новый слайс меньшей емкости
	result = append(result, arr[:i]...)   //копируем все элементы до i-го
	result = append(result, arr[i+1:]...) //копируем все элементы после i-го
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deleteElem(arr, 0))
	fmt.Println(deleteElem(arr, 5))
	fmt.Println(deleteElem(arr, 25)) //индекс за пределами массива
	fmt.Println(deleteElem(arr, -1)) //индекс отрицательный

}
