// За какое время выполнится скрипт
// Ответ: моментально
package main

import (
	"fmt"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("run go after 3 sec")
		ch <- 43
	}()

	return ch

}
func main() {
	//<-worker()
	//fmt.Println("After one")
	//<-worker()
	//fmt.Println("After two")
	// 6 секунд

	// =
	// 6 секунд, вызывает по очереди и присваивает
	// Сначала блочит выполнение на первом слушанье из канала, затем на втором
	_, _ = <-worker(), <-worker()

}
