package main

import (
	"fmt"
	"os"
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
			//при достижении i = 5 выходим из программы
			if i == 5 {
				fmt.Println("program exites with os.Exit")
				os.Exit(0)
			}
		}
	}()
	wg.Wait()
}
