package main

import (
	"fmt"
	"time"
)

// Дождаться выполнения части кода
func main() {
	waitForFinished()

}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		close(ch)
		fmt.Println("Sent signal")
	}()

	_, wd := <-ch
	fmt.Println("Receive signal", wd)

}
