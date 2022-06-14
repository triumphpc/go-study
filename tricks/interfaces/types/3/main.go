package main

import "fmt"

type A int

func (A) g()     {}
func (A) m() int { return 1 }

type B int

func (B) g() {}
func (B) f() {}

type C struct {
	A
	B
}

func (C) m() int {
	return 9
}

func main() {
	var c interface{} = C{}
	_, bf := c.(interface{ f() })
	_, bg := c.(interface{ g() })

	fmt.Println(bf) // true
	fmt.Println(bg) // false - при двойном embadding, не реализует этот интерфейс
	// Если идет эмбеддинг двух структур с реализациями одного и того же метода, то дочерняя структура не реализует этот интерфейс

	i, ok := c.(interface{ m() int })
	fmt.Println(ok, i.m()) // true 9
}
