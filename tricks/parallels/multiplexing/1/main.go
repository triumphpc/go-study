package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	tick := time.Tick(time.Second) // Тикает в канал раз в секунду
	abort := make(chan struct{})

	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			return
		}

		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		log.Println(countdown)

		select {
		case <-tick:
		case <-abort:
			log.Println("Race aborted")
			return
		}
	}

	start()
}

func start() {
	fmt.Println("Start race")
}
