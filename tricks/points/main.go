package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Race is starting... Press <enter> for stop")

	// abort listen
	abort := make(chan struct{})
	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			return
		}
		abort <- struct{}{}
	}()

	//tick := time.Tick(1 * time.Second) // тут утечка горутины
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	select {
	//	case <-tick:
	//		// Nothing,
	//	case <-abort:
	//		fmt.Println("Cancel!")
	//		return
	//	}
	//}

	ticker := time.NewTicker(2 * time.Second)
	select {
	case <-ticker.C:
		fmt.Println("tick")
	case <-abort:
		fmt.Println("Cancel!")
		return
	}
	ticker.Stop()

	start()

}

func start() {
	fmt.Println("Start race")
}
