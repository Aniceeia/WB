package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//закрываем контекст после выполнения всех встрок в main
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for i := 1; ; i++ {
			//при приходе сигнала done в контекст - завершаем программу
			select {
			case <-ctx.Done():
				fmt.Println("program ended with context timeout")
				return
				//если сигнал done не пришел - печетаем i
			case <-ticker.C:
				fmt.Println(i)
			}
		}
	}()
	//ставим main на паузу до завершения всех горутин
	wg.Wait()
}
