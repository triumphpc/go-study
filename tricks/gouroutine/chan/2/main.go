package main

import (
	"fmt"
	"time"
)

// Нельзя читать из закрытого небуферизированного канала
func main() {
	c := make(chan int)   // ошибка
	c = make(chan int, 1) // не будет ошибка
	go test(c)
	c <- 5

	time.Sleep(100 * time.Millisecond)

}

func test(c chan int) {
	time.Sleep(10 * time.Millisecond)
	close(c)
	fmt.Println(<-c)
}
