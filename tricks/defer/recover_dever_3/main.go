package main

import (
	"fmt"
)

var x int

func f() int {
	x++

	return x
}
func main() {
	o := fmt.Print

	defer o(f()) // 1. Обернуто в функцию. Высчитывает текущие значения, но не выполняет фукцию
	// x = 1
	// 5. Распечает 1

	defer func() {
		defer o(recover()) // 4. Восстанавливает работу и печатает 2

		o(f()) //  3. x=4 и печать
	}()

	defer f() // 2. x=3

	defer recover() // Прямой вызов из defer не восстанавливает работу, паника активна

	panic(f()) // 2. Вызывает панику и f(), x=2

}
