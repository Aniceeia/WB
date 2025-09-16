package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)            //канал для чисел
	stop := make(chan os.Signal, 1) //канал для сигнала остановки
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() { //запуск генератора чисел
		i := 0
		for {
			i++
			ch <- i
			time.Sleep(120 * time.Millisecond)
		}
	}()

	go func() { //горутина с таймером остановоки
		time.Sleep(5 * time.Second)
		fmt.Println("program stopped, you spend another 5 seconds")
		stop <- syscall.SIGINT //передаем сигнал остановки
	}()

	for {
		select {
		case res := <-ch: //получает числа из канала
			fmt.Println(res)
		case <-stop: //получает сигнал остановки
			close(ch)
			return
		}
	}

}
