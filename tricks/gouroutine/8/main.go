package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	//runtime.GOMAXPROCS(1)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" // может не успеть отработать, в случае одного процессора не выводет точно

	//time.Sleep(time.Second) // Если подождать, то успеет выполнить

	//close(ch) // Не обязательно закрывать
}
