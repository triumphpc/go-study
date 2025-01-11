package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1, ch2 chan int
	ch1 = make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	for {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
			ch1 = nil // Отключаем ch1 после получения значения
		case ch2 <- 2:
			fmt.Println("Sent to ch2")
		default:
			fmt.Println("No activity")
			time.Sleep(500 * time.Millisecond)
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
