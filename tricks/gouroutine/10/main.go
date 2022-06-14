// В канале nil операции отправки и приёма блокируются навсегда.
//Это хорошо задокументированное поведение, но оно может стать сюрпризом для новичков.
// Это поведение можно использовать для динамического включения и отключения блоков case в выражении select.

package main

import "fmt"
import "time"

func main() {
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int

		for {
			select {
			case out <- val: // (3) срабатывает конструкция и пишет в канал out значение
				fmt.Println("out ", val)
				out = nil
				in = inch
			case val = <-in: // (1) in это канал для чтения для inch
				fmt.Println("in", val)
				out = outch // (2) канал для записи инициализация
				in = nil    // блокируем для записи
			}
		}
	}()

	go func() {
		for r := range outch { // (4) считывает значения из канала
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
