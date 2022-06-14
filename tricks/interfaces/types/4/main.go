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
	f := t.M

	fmt.Printf("%T\n", f) // func()

	g := i.M              // Тут указатель слетает, так как указатель на интерфейс
	fmt.Printf("%T\n", g) // intfunc()
	t.X = 7

	fmt.Println(f(), g()) // 3 7
}
