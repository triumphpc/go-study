package main

import "fmt"

type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string // можно
	handler interface{}   // можно
	ref     *byte         // указатели можно
	raw     [10]byte      // слайсы можно
}

func main() {
	v1 := data{}
	v2 := data{}
	// Сравнение структур
	fmt.Println("v1 == v2:", v1 == v2) // выводит: v1 == v2: true
}
