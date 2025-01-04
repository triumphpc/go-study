package main

import "fmt"

//На каждом этапе горутины
//* получают значения от восходящего потока через входящие каналы,
//* выполняют некоторую функцию с этими данными, обычно производя новые значения,
//* отправляя значения вниз по потоку через исходящие каналы.

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}
