package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processInterface1(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		ch <- fmt.Sprintf("[coba1 coba2 coba3] %d", i+1)
		time.Sleep(time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)) * time.Millisecond)
	}
}

func processInterface2(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		ch <- fmt.Sprintf("[bisa1 bisa2 bisa3] %d", i+1)
		time.Sleep(time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 1; i++ {
		wg.Add(2)
		go processInterface1(&wg, ch)
		go processInterface2(&wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
