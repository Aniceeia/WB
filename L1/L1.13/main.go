package main

import "fmt"

func xOrSwap(a []int, first int, second int) {
	a[first] = a[first] ^ a[second]  //0001 (1) xor 0011 (3) = 0010 (2)
	a[second] = a[second] ^ a[first] //0011 (3) xor 0010 (2) = 0001 (1)
	a[first] = a[first] ^ a[second]  //0010 (2) (xor 0001 (1) = 0011 (3)
}

func main() {
	arr := []int{1, 2, 3} //0001 0010 0011
	xOrSwap(arr, 0, 2)
	fmt.Println(arr)
}
