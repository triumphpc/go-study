package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("1")
		wg.Done()
	}()

	wg.Wait() // Ожидание выполнение горутины

	go func() {
		fmt.Println(2)
	}()

	fmt.Println(3)

	// Будет выводить 1,  2 или 3
}
