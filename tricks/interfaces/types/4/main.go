package main

import "fmt"

type S struct {
	X int
}

func (s S) M() int {
	return s.X
}

type T struct {
	S
}

type I interface {
	M() int
}

func main() {
	t := &T{S{3}}

	var i I = t
	f := t.M // Тут идет замыкания и берется текущее значение структуры

	fmt.Printf("%T\n", f) // func() int

	g := i.M              // Тут указатель на тип как интерфейс идет
	fmt.Printf("%T\n", g) // func() int

	t.X = 7

	fmt.Println(f(), g()) // 3 7

	f2 := t.M // А тут уже обновленное значение

	fmt.Println(f2(), g()) // 7 7
}
