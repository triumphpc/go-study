package main

import (
	"fmt"
)

// Позволяет избежать утечки горутины по done каналу

func main() {

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	done := make(chan interface{})
	myChan := make(chan interface{})

	go func() {
		defer close(done)
		for i := 0; i < 5; i++ {
			myChan <- i
		}

	}()

	for val := range orDone(done, myChan) { // Do something with val
		fmt.Println(val) // Только до 3х так как done закрое первым
	}
	defer close(myChan)

}
