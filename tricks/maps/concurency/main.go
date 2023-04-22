// Что выведет код?
package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	fmt.Println(len(m)) // 0

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(N)
	for i := 0; i <= N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	fmt.Println(len(m)) // 10

	wg.Wait()

	//mu.Lock()
	fmt.Println(m) // map[0:0 1:1 3:3 4:4 5:5 6:6 7:7 8:8 9:9 10:10]
	//mu.Unlock()

}
