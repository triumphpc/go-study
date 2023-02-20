package main

// Аргументы для вызовов отложенных функций вычисляются тогда же, когда и выражение
//defer (а не когда на самом деле выполняется функция).

import "fmt"

func main() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }()) // 2

	defer func() { // функция выполняется и берет значения выходящих уже параметров
		fmt.Println("result =>", func() int { return i * 2 }()) // 6
	}()
	i++

	defer fmt.Println("result =>", func() int { return i * 2 }()) // 4
	i++
}
