package main

import "fmt"

func main() {
	// Не выделяет память, если не используется слайс
	var a []int
	fmt.Printf("%T %v %p\n", a, a, a)

	// Выделяем память и устанавливает nil значение
	b := []int{}
	fmt.Printf("%T %v %p", b, b, b)

}
