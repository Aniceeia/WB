package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val uint64
}

func (c *Counter) increment() {
	atomic.AddUint64(&c.val, 1)
}

func (c *Counter) getValue() uint64 {
	return atomic.LoadUint64(&c.val)
}

func main() {
	var wg sync.WaitGroup //синзронизируем main

	c := &Counter{} //инициализация counter

	for range 1000 { //создали 1000 горутин
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment() //прибавляем 1 с помощью атомик
		}()
	}
	wg.Wait()
	fmt.Println("maximum value", c.getValue()) // выводим максимальное значение counter
}
