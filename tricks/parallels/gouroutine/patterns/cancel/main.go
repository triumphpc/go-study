package main

import (
	"fmt"
	"time"
)

// Cancel - Иногда нужно указать go-подпрограмме, что она должна прекратить свои дей­ствия, например, на веб-сервере, где она выполняет вычисления для клиента, соеди­нение с которым разорвано.
func main() {
	doWork := func(strings <-chan string, done <-chan struct{}) <-chan struct{} {
		completed := make(chan struct{})

		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(completed)
			}()

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}

			}

		}()
		fmt.Println("exit")
		return completed
	}

	st := make(chan string)
	done := make(chan struct{})
	terminate := doWork(st, done)

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Cancel")
		close(done)
	}()

	<-terminate // При закрытии канала происходит разблокировка горутины
	fmt.Println("Done")
	//exit
	//test
	//doWork exited.

}
