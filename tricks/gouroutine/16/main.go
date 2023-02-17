package main

import "fmt"

func main() {
	//c := make(chan string)

	//go func() {
	//	fmt.Println(<-c) // если обернут в функцию и вызывать первым объявлением, то проблем не будет
	//}()

	//c <- "write data" // тут будет deadlock, потому что небуферизированный канал лочит
	// основной стек.

	//go fmt.Println(<-c) // а вот так будет deadlock

	//fmt.Println("after write")
	//close(c)

	// Для буферизированных работает иначе

	ch := make(chan bool, 4)
	ch <- true
	ch <- false

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) /// тут будет дедлок стека, так как нет данных

	close(ch)
}
