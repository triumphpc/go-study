package main

import (
	"fmt"
	"time"
)

// Cancel - Иногда нужно указать go-подпрограмме, что она должна прекратить свои дей­ствия, например, на веб-сервере, где она выполняет вычисления для клиента, соеди­нение с которым разорвано.

func main() {
	doWork := func(strings <-chan int) <-chan struct{} {
		completed := make(chan struct{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(completed)
			}()

			for s := range strings {
				fmt.Println(s)
			}
		}()
		fmt.Println("exit")
		return completed
	}

	st := make(chan int)
	terminate := doWork(st)

	go func() {
		for i := 0; i < 3; i++ {
			st <- i
			time.Sleep(time.Second * 1)
		}

		close(st)

	}()

	<-terminate
	fmt.Println("Done")

}
