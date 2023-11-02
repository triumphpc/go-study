package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("empl send signal")
	}()

	tc := time.After(1 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("receive p signal", p)
	case t := <-tc:
		fmt.Println("receive t signal", t) // Если горутина выйдет по этиму сигналу, то будет утечка горутины с каналом ch
	}

	time.Sleep(time.Second)
	fmt.Println("------------------")

}
