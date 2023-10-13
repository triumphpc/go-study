package main

import (
	"fmt"
	"math/rand"
	"time"
)

// * Чтение из канала <-c блокирует горутину навсегда
//* Отправка в канал c <- v блокирует горутину навсегда
//* Закрытие канала close(c) вызовет панику.

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Millisecond * 100)
	for {
		select {
		case input := <-c:
			fmt.Println("Читаю")
			sum = sum + input
		case <-t.C:
			c = nil // Блокирует горутину навсегда на чтение и на запись
			fmt.Println(sum)

			fmt.Println("После этого ничего уже не генерится ")
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
		fmt.Println("Пишу")
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)
	time.Sleep(time.Second)

}
