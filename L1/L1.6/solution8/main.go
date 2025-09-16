package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
			//при достижении 500 милисекунд или 5 итерации - останавливаем программу
			if i == 5 {
				runtime.Goexit()
			}
		}
	}()

	wg.Wait()
	fmt.Println("exit with goexit")
}
