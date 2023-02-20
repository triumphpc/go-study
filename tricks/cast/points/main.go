package main

import "fmt"

type myInt int
type intPtr *int
type MyIntPtr *myInt

func convert(x MyIntPtr) (y intPtr) {
	// Преобразование разных типов ссылочных

	b := int(*x)
	y = &b

	return y
}

func main() {
	var a myInt

	a = 5
	b := convert(&a)

	fmt.Printf("%#v", *b)
}
