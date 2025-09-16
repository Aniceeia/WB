package main

import (
	"fmt"
	"time"
)

func main() {
	//остановить данную программу можно только через ctrl+c
	//ungracefull shutdown
	ch := make(chan int)

	go func() {
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}

}
