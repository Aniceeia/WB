package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//создаем момент времени на 1с больше чем настоящее время
	deadline := time.Now().Add(time.Second)
	//создаем контекст с функцией завершения ко времени
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	//создаем переменные для хранения счетчика синхронизации
	var wg sync.WaitGroup
	wg.Add(1)
	//запускаем горунтину имитирующую как-либо процесс
	go func() {
		defer wg.Done()
		//так как процесс периодичный используем тикерц
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for i := 1; ; i++ {
			select {
			//если прошло 100 мс выводим i
			case <-ticker.C:
				fmt.Println(i)
				//если из контекста пришел сигнал завершения разрываем цикл
			case <-ctx.Done():
				fmt.Println("program ended with context deadline")
				return
			}
		}
	}()
	//ждем завершения вызавнной горутины в замыкании
	wg.Wait()
}
