package main

import (
	"fmt"
	"sync"
)

func main() {
	temps := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}
	groups := make(map[int][]float64)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, v := range temps {
		wg.Add(1)
		go func(v float64) {
			defer wg.Done()
			category := int(v/10) * 10
			mu.Lock()
			groups[category] = append(groups[category], v)
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	fmt.Println(groups)
}
