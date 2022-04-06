package main

import "fmt"

type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

func main() {
	v1 := data{}
	v2 := data{}
	// Сравнение структур
	fmt.Println("v1 == v2:", v1 == v2) // выводит: v1 == v2: true
}
