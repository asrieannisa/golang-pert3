package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func processInterfaceMutex1(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		mu.Lock()
		ch <- fmt.Sprintf("[coba1 coba2 coba3] %d", i+1)
		mu.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func processInterfaceMutex2(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		mu.Lock()
		ch <- fmt.Sprintf("[bisa1 bisa2 bisa3] %d", i+1)
		mu.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 1; i++ {
		wg.Add(2)
		go processInterfaceMutex1(&wg, ch)
		go processInterfaceMutex2(&wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
