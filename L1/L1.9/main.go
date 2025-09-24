package main

import "fmt"

func double(ch chan int, done chan struct{}) chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for v := range ch {

			select {
			case <-done:
				return
			case result <- v * 2:
			}
		}
	}()
	return result
}

func writeToChan(arr []int, done chan struct{}) chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		for _, v := range arr {
			select {
			case <-done:
				return
			case result <- v:
			}
		}
	}()
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	done := make(chan struct{}) //канал для остановки работы горутин
	defer close(done)
	ch := writeToChan(arr, done) //записываем множество в канал
	resultCh := double(ch, done) //передаем канал в функцию удвоения
	for v := range resultCh {
		fmt.Println(v) //читаем из канала
	}
}
