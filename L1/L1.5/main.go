package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	N := 2
	timeout := time.After(time.Duration(N) * time.Second)

	go func() {
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(time.Second)

		}
	}()

	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			return
		case res := <-ch:
			fmt.Println(res)
		}
	}
}
