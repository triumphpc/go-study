package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock() // Процесс 1 блокирует А; Процесс 2 блокирует B
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock() // Процесс 1 ожидает B; Процесс 2 ожидает А
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // Процесс 1
	go printSum(&b, &a) // Процесс 2
	wg.Wait()
}

type value struct {
	mu    sync.Mutex
	value int
}
