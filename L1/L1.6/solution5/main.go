package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//создаем когтекст в условием выхода - cancel
	ctx, cancel := context.WithCancel(context.Background())

	go func() { //создаем горутину для оработки ввода
		defer cancel()
		fmt.Scanln()
		fmt.Println("manual cancel with enter key")
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	//запуск рабочей горутины
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for i := 1; ; i++ {
			select {
			case <-ticker.C:
				fmt.Println(i)
			case <-ctx.Done():
				return
			}

		}
	}()
	//ожидаем завершения в главной горутине
	wg.Wait()
	fmt.Println(ctx.Err())
}
