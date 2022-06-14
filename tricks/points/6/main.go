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

	fmt.Printf("%+v %+p\n", p, p)
	fmt.Printf("%+v %+p\n", p1, p1)
}
