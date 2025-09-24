package main

import (
	"fmt"
	"sync"
)

func main() {
	temps := []float64{
		-25.4, -27.0, -21.0, -29.9,

		-19.9, -20.0, -20.1,

		-5.2, -3.7, -0.5, -9.9,

		0.0, 0.1, 5.7, 9.9,

		13.0, 19.0, 15.5, 24.5,

		19.9, 20.0, 20.1, 29.9,

		32.5, 35.8, 40.2,

		-40.1, 45.9, 50.0,

		9.999, 10.0001, 19.999, 20.0001,

		-10.001, -9.999, -0.001,
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
