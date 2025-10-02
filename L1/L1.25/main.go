package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	if duration <= 0 {
		return
	}

	ch := make(chan struct{}) //создаем канал структуры, для сигнала

	go func() {
		timer := time.NewTimer(duration) //создаем таймер
		<-timer.C                        //ждем срабатывания данных
		close(ch)                        // закрываем канал
	}()
	<-ch //блокируем функции до закрытия канала
}

func main() {
	fmt.Println(time.Now().Format("15:04:05"))
	Sleep(2 * time.Second)
	fmt.Println(time.Now().Format("15:04:05"))
}
