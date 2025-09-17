package main

import (
	"fmt"
	"sync"
)

// так как данные необходимо в первую очерель записать, а не прочитать, создадим свою структуру
type mapWrite struct {
	m  map[int]int
	mu sync.RWMutex
}

func main() {
	m := mapWrite{
		m: make(map[int]int),
	}

	var wg sync.WaitGroup //создает waitgroup чтобы убедистя что все горутины заверщены

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.mu.Lock()
			m.m[i]++
			m.mu.Unlock()
		}(i)
	}
	//ждем завершения всех горутин
	wg.Wait()

	m.mu.RLock() //читаем тоже с блокировкой
	defer m.mu.RUnlock()

	for v := range m.m {
		fmt.Println(v) //результат конкуретной записи
	}
	fmt.Println("end of map")
}
