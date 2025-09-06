package main

import (
	"fmt"
	"sync"
)

func square(idx int, n int, res []int, wg *sync.WaitGroup) {
	defer wg.Done()
	res[idx] = n * n
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := make([]int, len(arr))
	var wg sync.WaitGroup

	for i, val := range arr {
		wg.Add(1)
		go square(i, val, res, &wg)

	}
	wg.Wait()
	fmt.Printf("%v", res)
}
