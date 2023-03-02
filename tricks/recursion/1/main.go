//Программист Игорь решил добавить модуль для парсинга данных в стандартную
//библиотеку своего языка. Одним из модулей парсера будет конечный автомат,
//войдя в поток, Игорь набросал следующий код, что произойдёт в момент первой компиляции?

package main

import "fmt"

type state func(x int) state

func start(x int) state {
	fmt.Println("start", x)
	if x == 0 {
		return middle
	} else {
		return end
	}
}

func middle(_ int) state {
	fmt.Println("end")

	return end
}

func end(_ int) state {
	fmt.Println("start")

	return start
}

func main() {
	state := start(0)
	state = state(1)
	state = state(2)
	state = state(0)
}
