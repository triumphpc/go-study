// Какая из строк (A или B) позволит получить лучшее время выполнения при условии что у нас 8 ядер (numcpu=8)?
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU:", numcpu)

	runtime.GOMAXPROCS(1) ////Опять же это быстрее, потому что меньше блокировок
	//runtime.GOMAXPROCS(numcpu)

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() {
		for i := 0; i < 1000000; i++ {
			ch1 <- i
		}

		fmt.Println("here")

		ch1 <- -1
		ch2 <- 0.0
	}()

	go func() {
		total := 0.0

		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < 100000; i++ {
				m := <-ch1

				if m == -1 {
					fmt.Println("here 2")
					fmt.Println("Total:", total)
					ch2 <- total
				}
			}

			t2 := time.Now().UnixNano()
			dt := float64(t2-t1) / 1000000.0
			total += dt
			fmt.Println(dt)
		}
	}()

	time.Sleep(time.Second * 3)

	fmt.Println(<-ch2)

}
