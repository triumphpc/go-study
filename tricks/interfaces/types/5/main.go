package main

import "log"

// Как узнать типы интерфейса ?

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

	switch i.(type) {
	case T:
		log.Println("T YES")
	case I:
		log.Println("YES") // YES

	}
}
