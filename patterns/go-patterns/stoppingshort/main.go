package main

import (
	"fmt"
	"sync"
)

// Этот шаблон позволяет записать каждую стадию приема в виде цикла диапазона и гарантирует,
//что все горутины завершатся после того, как все значения будут успешно отправлены вниз по течению.
// https://go.dev/blog/pipelines

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums)) // Here buffering channel
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1) // enough space for the unread inputs

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from the output.
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Since we didn't receive the second value from out,
	// one of the output goroutines is hung attempting to send it.

	// We need to arrange for the upstream stages of our pipeline to exit even when the downstream stages
	//fail to receive all the inbound values. One way to do this is to change the outbound channels to have a buffer.
	//A buffer can hold a fixed number of values; send operations complete immediately if there’s room in the buffer:

	//c := make(chan int, 2) // buffer size 2
	//c <- 1  // succeeds immediately
	//c <- 2  // succeeds immediately
	//c <- 3  // blocks until another goroutine does <-c and receives 1

}
