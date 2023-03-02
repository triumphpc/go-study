package main

import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{"hello"}
}

func main() {
	p := f()
	p1 := *f()

	fmt.Printf("%+v %p\n", p, p) // &{m:hello} 0xc000096210 указатель и адрес

	fmt.Printf("%+v %p\n", p1, p1) // {m:hello} %!p(main.S={hello}) // значение, на какой тип ссылается и какое значение

}
