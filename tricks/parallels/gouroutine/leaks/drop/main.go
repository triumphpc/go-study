package main

import "fmt"

// Для случая, когда мы пишем в буфферизированный канал данные, если в момент записи буффер переполнен, мы можем “бросать” данные, использую конструкцию select/case
//В этом случае горутина не блокирует выполнение программы.

func main() {
	drop()

}
func drop() {
	const cap = 5

	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("recv'd signal", p)
		}
	}()

	const work = 20

	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("Sent", w)
		default:
			fmt.Println("Dropped")
		}
	}

	close(ch)

}
