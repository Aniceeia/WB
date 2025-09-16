package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)                   //канал для чисел
	timeout := time.After(3 * time.Second) //таймер на 3 сек

	go func() {
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(250 * time.Millisecond)
		}
	}()

	for {
		select {
		case <-timeout: //срабатывает через 3 сек
			fmt.Println("timeout, you spend 3 second")
			return
		case res := <-ch: //читает числа из канала
			fmt.Println(res)
		}
	}
}
