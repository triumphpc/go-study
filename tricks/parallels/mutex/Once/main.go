package main

// Если в группе горутин нужно убедиться, что функция выполнится один раз - используем once

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	increment := func() {
		count++
	}
	var once sync.Once
	var increments sync.WaitGroup

	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count) // Count is 1

}
