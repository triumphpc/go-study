package main

import "fmt"

func main() {
	c := make(chan string)

	c <- "write data" // тут будет deadlock, потому что небуферизированный канал лочит
	// основной стек.

	//go func() {
	//	fmt.Println(<-c) // если обернут в функцию и вызывать первым объявлением, то проблем не будет
	//}()

	go fmt.Println(<-c) // а вот так будет deadlock

	fmt.Println("after write")
}
