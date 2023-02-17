// Пример с багом можно исправить, сигнализируя через специальный канал отмены (special cancellation channel)
//остальным рабочим горутинам, что их результаты больше не нужны.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "sent result")
			case <-done:
				fmt.Println(idx, "exiting")
			}
			fmt.Println(idx, "here")
		}(i)
	}

	// get first result
	fmt.Println("result:", <-ch)
	fmt.Println("result:", <-ch)
	fmt.Println("result:", <-ch)
	close(done)
	// do other work
	time.Sleep(3 * time.Second)
}
