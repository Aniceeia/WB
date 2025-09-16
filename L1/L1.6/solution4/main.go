package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//создаем очередь сообщений
	ch := make(chan int)
	var wg sync.WaitGroup
	//запускаем потребителя
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range ch { //на этой операции горутина тормозит
			fmt.Println(data)
		}
	}()
	//запускаем производителя
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) //пауза между отправками
	}

	close(ch) //закрывает канал, сообщает потребителю что данных больше не будет3
	wg.Wait() //ожидает завершения записи в потребитель
	fmt.Println("closed channel")
}
