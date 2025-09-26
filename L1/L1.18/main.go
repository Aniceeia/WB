package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) getValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.getValue())

}
