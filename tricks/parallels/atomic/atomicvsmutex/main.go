package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// BenchmarkBasic-16     	1000000000	         0.01186 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBasic2-16    	1000000000	         0.01411 ns/op	       0 B/op	       0 allocs/op

// Атомарные операции быстрее работают чем мютексы

func f(v *uint32, wg *sync.WaitGroup) {
	for i := 0; i < 300000; i++ {
		atomic.AddUint32(v, 1)
		//*v++
	}
	wg.Done()
}

func f2(v *uint32, wg *sync.WaitGroup, mx *sync.Mutex) {
	for i := 0; i < 300000; i++ {
		mx.Lock()
		*v++
		mx.Unlock()
	}
	wg.Done()
}

func main() {
	var v uint32 = 42
	var wg sync.WaitGroup
	wg.Add(2)
	mx := sync.Mutex{}

	go f2(&v, &wg, &mx)
	go f2(&v, &wg, &mx)
	wg.Wait()

	fmt.Println(v)
}
