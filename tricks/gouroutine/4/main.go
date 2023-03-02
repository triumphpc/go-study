// Что выведет данный код?
// Должен войти в бесконечный цикл
// Но так как один тред может обрабатывать две горутины, то планировщик может успеть обработать и выйти из цикла
package main

import (
	"fmt"
	"runtime"
)

func main() {
	val := runtime.GOMAXPROCS(1)
	fmt.Println(val) // Возвращает текущее значение = 16

	val = runtime.GOMAXPROCS(0)
	fmt.Println(val) // 1

	val = runtime.GOMAXPROCS(0)
	fmt.Println(val) // 1

	done := false

	fmt.Println("run")

	go func() {
		fmt.Println("go 1")
		done = true
	}()

	fmt.Println("go 2")
	for !done {
		fmt.Println("x")
	}

	fmt.Printf("Done")
}
