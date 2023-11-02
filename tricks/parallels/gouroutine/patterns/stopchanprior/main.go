package main

import (
	"fmt"
	"time"
)

// Приоритезация стоп канала

func main() {

	ch := make(chan struct{})
	chMsg := make(chan string)

	//go f(ch)
	f2(ch, chMsg)

	go func() {
		for {
			chMsg <- "ttt"
			ch <- struct{}{}
		}

	}()

}

func f2(ch chan struct{}, chMsg chan string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ch:
			fmt.Println("exit")
			return
		default:
		}

		select {
		case <-ticker.C:
			fmt.Println("some work")
			return
		default:
		}

		select {
		case <-ticker.C: // Всегда будет этот отрабатывать, так как выделен в отдельный блок
			fmt.Println("some work")
		case <-ch:
			fmt.Println("exit")
			return
		case msg := <-chMsg:
			fmt.Println(msg)
			return
		}
	}
}
