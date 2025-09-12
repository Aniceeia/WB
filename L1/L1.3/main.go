package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	exit = 'q'
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{})

	var wg sync.WaitGroup

	workers, err := inputWorkers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go workerLoop(i, ch, &wg)
	}

	go stopSignal(stop)

	writeToChannel(ch, stop, &wg)

}

func workerLoop(worker int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for c := range ch {
		fmt.Printf("worker: %d, doing %d\n", worker, c)
	}
}

func stopSignal(stop chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			continue
		}
		if char == exit {
			close(stop)
			return
		}

	}
}

func inputWorkers() (int, error) {
	if len(os.Args) < 2 {
		return 0, fmt.Errorf("put argument when compile")
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		return 0, fmt.Errorf("invalid number")
	}
	return workers, nil
}

func writeToChannel(ch chan int, stop <-chan struct{}, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-stop:
			close(ch)
			wg.Wait()
			fmt.Println("program stopped")
			return
		default:
			ch <- i
			time.Sleep(time.Second)
		}

	}
}
