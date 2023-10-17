package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	c := make(chan int)
	go test(c)
	c <- 5

	//c <- 5 // Ошибка записи в закрытый канал

	c = make(chan int, 2)
	go test2(c)
	c <- 5
	c <- 5

	time.Sleep(10 * time.Millisecond)
	//c <- 5 // Ошибка,нельзя писать в закрытй

	c = make(chan int, 2)
	go test3(c)

	time.Sleep(10 * time.Millisecond)
	c <- 55

	time.Sleep(100 * time.Millisecond)
}

func test(c chan int) {
	fmt.Println(<-c) // 5
	close(c)
	fmt.Println(<-c) // 0
}

func test2(c chan int) {
	close(c)
	fmt.Println(<-c) // 5
	fmt.Println(<-c) // 5
	fmt.Println(<-c) // 0

}

func test3(c chan int) {
	log.Println("Старт")
	fmt.Printf("`%#v`\n", <-c) // Тут блокируется пока не прочитает
	log.Println("Выход")
}
