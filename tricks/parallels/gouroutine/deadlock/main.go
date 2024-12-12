package main

import "fmt"

func main() {
	//c := make(chan string)
	//
	//go func() {
	//	fmt.Println(<-c) // если обернут в функцию и вызывать первым объявлением, то проблем не будет
	//}()
	//
	//c <- "write data"
	//
	//close(c)

	// Пример 1. Deadlock из-за блокировки не буферизированного канал
	c := make(chan string)
	c <- "write data" // тут будет deadlock, потому что небуферизированный канал лочит на чтение основной стек.

	go fmt.Println(<-c) // а вот так будет deadlock, аналогично

	fmt.Println("after write")
	close(c)

	// Пример 2. Для буферизированных работает иначе
	//ch := make(chan bool, 4)
	//ch <- true
	//ch <- false
	//
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	////fmt.Println(<-ch) // тут будет дедлок стека, так как нет данных
	//
	//close(ch)
	// Пример 3. Без дедлока

	//ch = make(chan bool, 4)
	//ch <- true
	//ch <- false
	//
	//go func() {
	//	fmt.Println(<-ch)
	//	fmt.Println(<-ch)
	//	log.Println("Start")
	//	fmt.Println(<-ch) // тут лочится до момента закрытия канала
	//	log.Println("Stop")
	//}()
	//time.Sleep(10 * time.Millisecond)
	//
	//close(ch)
	//
	//time.Sleep(10 * time.Millisecond)

}
