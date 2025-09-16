package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var stop int32 = 0 //stop 0(false)

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for atomic.LoadInt32(&stop) == 0 {
			i++
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
	}()
	time.Sleep(500 * time.Millisecond)
	atomic.StoreInt32(&stop, 1) //устанавливаем stop = 1(true)
	wg.Wait()
	fmt.Println("ended with atomic int32")
}
