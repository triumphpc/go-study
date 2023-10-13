package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		select {
		case <-time.Tick(time.Second * 5):
			fmt.Println("fdf")
		}
	}
}
